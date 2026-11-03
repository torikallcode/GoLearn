package service

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"golearn/internal/model"
	"golearn/internal/runner"
	"golearn/internal/store"
)

type LearnService struct {
	store  *store.Store
	runner runner.Runner
	now    func() time.Time
}

func NewLearnService(s *store.Store, r runner.Runner) *LearnService {
	return &LearnService{store: s, runner: r, now: time.Now}
}

func (ls *LearnService) Run(ctx context.Context, code string) (runner.Result, error) {
	if strings.TrimSpace(code) == "" {
		return runner.Result{}, NewError(http.StatusUnprocessableEntity, "validation", "Kode tidak boleh kosong.")
	}
	res, err := ls.runner.Run(ctx, code, "")
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return runner.Result{}, NewError(http.StatusGatewayTimeout, "timeout", "Waktu eksekusi habis.")
		}
		return runner.Result{}, ErrRunner
	}
	return res, nil
}

type TestResult struct {
	Index    int    `json:"index"`
	Hidden   bool   `json:"hidden"`
	Passed   bool   `json:"passed"`
	Expected string `json:"expected,omitempty"`
	Actual   string `json:"actual"`
	Stderr   string `json:"stderr,omitempty"`
}

type SubmissionResult struct {
	Passed   bool         `json:"passed"`
	Tests    []TestResult `json:"tests"`
	XPEarned int          `json:"xp_earned"`
	Stats    model.Stats  `json:"stats"`
	Message  string       `json:"message"`
}

func (ls *LearnService) SubmitExercise(ctx context.Context, userID, exerciseID uint64, code string) (*SubmissionResult, error) {
	if strings.TrimSpace(code) == "" {
		return nil, NewError(http.StatusUnprocessableEntity, "validation", "Kode tidak boleh kosong.")
	}
	ex, err := ls.store.ExerciseByID(exerciseID)
	if err != nil {
		return nil, mapStoreErr(err)
	}
	tests, err := ls.store.ExerciseTests(exerciseID)
	if err != nil {
		return nil, err
	}
	if len(tests) == 0 {
		return nil, NewError(http.StatusUnprocessableEntity, "no_tests", "Latihan ini belum punya test case.")
	}
	if len(tests) > 10 {
		tests = tests[:10]
	}
	out := &SubmissionResult{Tests: make([]TestResult, 0, len(tests)), Passed: true}
	var lastOut strings.Builder
	for i, tc := range tests {
		if tc.Input != "" && !ls.runner.SupportsStdin() {
			return nil, NewError(http.StatusServiceUnavailable, "runner_stdin",
				"Latihan ini butuh input. Jalankan server dengan RUNNER_MODE=docker.")
		}
		res, runErr := ls.runner.Run(ctx, code, tc.Input)
		if runErr != nil {
			return nil, ErrRunner
		}
		passed := res.Exit == 0 && normalizeOutput(res.Stdout) == normalizeOutput(tc.ExpectedOutput)
		tr := TestResult{Index: i + 1, Hidden: tc.IsHidden, Passed: passed, Actual: res.Stdout, Stderr: res.Stderr}
		if !tc.IsHidden {
			tr.Expected = tc.ExpectedOutput
		}
		if !passed {
			out.Passed = false
		}
		out.Tests = append(out.Tests, tr)
		lastOut.WriteString(res.Stdout)
		lastOut.WriteString(res.Stderr)
	}

	if err = ls.store.CreateSubmission(userID, exerciseID, code, out.Passed, lastOut.String()); err != nil {
		return nil, err
	}
	if rErr := ls.store.MarkLessonStarted(userID, ex.LessonID); rErr != nil {
		return nil, rErr
	}
	if out.Passed {
		prior, err := ls.store.PriorPassedSubmission(userID, exerciseID)
		if err != nil {
			return nil, err
		}
		if !prior {
			out.XPEarned = ex.XPReward
			out.Stats, err = ls.award(userID, ex.XPReward)
			if err != nil {
				return nil, err
			}
			out.Message = "Semua test lolos. Kerja bagus!"
		} else {
			out.Stats, _ = ls.store.GetStats(userID)
			out.Message = "Semua test lolos."
		}
	} else {
		out.Stats, _ = ls.store.GetStats(userID)
		out.Message = "Belum semua test lolos. Cek output lalu coba lagi."
	}
	return out, nil
}

func normalizeOutput(s string) string {
	return strings.TrimSpace(strings.ReplaceAll(s, "\r\n", "\n"))
}

type AnswerResult struct {
	Correct         bool        `json:"correct"`
	CorrectOptionID *uint64     `json:"correct_option_id,omitempty"`
	CorrectAnswer   string      `json:"correct_answer,omitempty"`
	Explanation     string      `json:"explanation"`
	XPEarned        int         `json:"xp_earned"`
	Stats           model.Stats `json:"stats"`
}

func (ls *LearnService) AnswerQuiz(userID, quizID uint64, optionID *uint64, answer string) (*AnswerResult, error) {
	q, err := ls.store.QuizByID(quizID)
	if err != nil {
		return nil, mapStoreErr(err)
	}
	out := &AnswerResult{Explanation: q.Explanation}
	correct := false
	if q.Type == "short" {
		if q.CorrectAnswer != nil && strings.EqualFold(strings.TrimSpace(answer), strings.TrimSpace(*q.CorrectAnswer)) {
			correct = true
		}
		if q.CorrectAnswer != nil {
			out.CorrectAnswer = *q.CorrectAnswer
		}
	} else {
		for _, o := range q.Options {
			if o.Correct {
				id := o.ID
				out.CorrectOptionID = &id
			}
			if optionID != nil && o.ID == *optionID && o.Correct {
				correct = true
			}
		}
	}
	out.Correct = correct
	prior, err := ls.store.HasCorrectAttempt(userID, quizID)
	if err != nil {
		return nil, err
	}
	if err = ls.store.CreateQuizAttempt(userID, quizID, correct); err != nil {
		return nil, err
	}
	if err = ls.store.MarkLessonStarted(userID, q.LessonID); err != nil {
		return nil, err
	}
	if correct && !prior {
		out.XPEarned = 5
		if out.Stats, err = ls.award(userID, 5); err != nil {
			return nil, err
		}
	} else if out.Stats, err = ls.store.GetStats(userID); err != nil {
		return nil, err
	}
	return out, nil
}

type CompleteResult struct {
	AlreadyCompleted bool        `json:"already_completed"`
	Stats            model.Stats `json:"stats"`
	DueReviews       int         `json:"due_reviews"`
}

func (ls *LearnService) CompleteLesson(userID uint64, slug string) (*CompleteResult, error) {
	l, _, p, err := ls.store.LessonBySlug(slug)
	if err != nil {
		return nil, mapStoreErr(err)
	}
	all, err := ls.store.PathLessonList(p.ID)
	if err != nil {
		return nil, err
	}
	completed, err := ls.store.CompletedLessonIDs(userID)
	if err != nil {
		return nil, err
	}
	if !completed[l.ID] && lessonStatuses(all, completed)[l.ID] == "locked" {
		return nil, ErrLocked
	}
	now := ls.now()
	changed, err := ls.store.CompleteLesson(userID, l.ID, now)
	if err != nil {
		return nil, err
	}
	out := &CompleteResult{AlreadyCompleted: !changed}
	if changed {
		if err = ls.store.SeedReviews(userID, l.ID, now); err != nil {
			return nil, err
		}
		out.Stats, err = ls.award(userID, 20)
		if err != nil {
			return nil, err
		}
	} else {
		out.Stats, err = ls.store.GetStats(userID)
		if err != nil {
			return nil, err
		}
	}
	out.DueReviews, _ = ls.store.DueCardCount(userID, now)
	return out, nil
}

func (ls *LearnService) award(userID uint64, xp int) (model.Stats, error) {
	st, err := ls.store.Award(userID, xp, ls.now())
	if err != nil {
		return st, err
	}
	total, _ := ls.store.TotalCompleted(userID)
	if total >= 1 {
		_ = ls.store.AwardBadge(userID, "first-lesson")
	}
	if st.CurrentStreak >= 7 {
		_ = ls.store.AwardBadge(userID, "streak-7")
	}
	if st.CurrentStreak >= 30 {
		_ = ls.store.AwardBadge(userID, "streak-30")
	}
	return st, nil
}
