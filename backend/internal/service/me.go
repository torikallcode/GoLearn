package service

import (
	"net/http"
	"time"

	"golearn/internal/model"
	"golearn/internal/store"
)

type MeService struct {
	store *store.Store
	now   func() time.Time
}

func NewMeService(s *store.Store) *MeService {
	return &MeService{store: s, now: time.Now}
}

func (m *MeService) Dashboard(userID uint64) (*model.Dashboard, error) {
	now := m.now()
	stats, err := m.store.GetStats(userID)
	if err != nil {
		return nil, err
	}
	d := &model.Dashboard{Stats: stats}
	if d.DueReviews, err = m.store.DueCardCount(userID, now); err != nil {
		return nil, err
	}
	if d.CompletedToday, err = m.store.CompletedToday(userID, now); err != nil {
		return nil, err
	}
	if d.TotalCompleted, err = m.store.TotalCompleted(userID); err != nil {
		return nil, err
	}
	if d.TotalLessons, err = m.store.TotalPublishedLessons(); err != nil {
		return nil, err
	}
	if d.Recent, err = m.store.RecentProgress(userID, 5); err != nil {
		return nil, err
	}
	if d.Paths, err = m.store.ProgressByPath(userID); err != nil {
		return nil, err
	}
	d.Continue, err = m.continueLesson(userID)
	return d, err
}

func (m *MeService) continueLesson(userID uint64) (*model.ContinueLesson, error) {
	paths, err := m.store.Paths()
	if err != nil {
		return nil, err
	}
	completed, err := m.store.CompletedLessonIDs(userID)
	if err != nil {
		return nil, err
	}
	for _, p := range paths {
		lessons, err := m.store.PathLessonList(p.ID)
		if err != nil {
			return nil, err
		}
		for _, l := range lessons {
			if !completed[l.ID] {
				return &model.ContinueLesson{Slug: l.Slug, Title: l.Title, PathSlug: p.Slug, Minutes: l.EstimatedMinutes}, nil
			}
		}
	}
	return nil, nil
}

func (m *MeService) Progress(userID uint64) ([]model.PathProgress, error) {
	return m.store.ProgressByPath(userID)
}

func (m *MeService) Stats(userID uint64) (map[string]any, error) {
	stats, err := m.store.GetStats(userID)
	if err != nil {
		return nil, err
	}
	badges, err := m.store.Badges(userID)
	if err != nil {
		return nil, err
	}
	due, err := m.store.DueCardCount(userID, m.now())
	if err != nil {
		return nil, err
	}
	total, _ := m.store.TotalCompleted(userID)
	return map[string]any{
		"stats":            stats,
		"badges":           badges,
		"due_reviews":      due,
		"lessons_done":     total,
		"xp_to_next_level": 100 - stats.XP%100,
	}, nil
}

func (m *MeService) ReviewToday(userID uint64) ([]model.DueCard, error) {
	return m.store.DueCards(userID, m.now(), 20)
}

type ReviewResult struct {
	NextReviewAt time.Time `json:"next_review_at"`
	IntervalDays int       `json:"interval_days"`
	Repetitions  int       `json:"repetitions"`
	EaseFactor   float64   `json:"ease_factor"`
	XPEarned     int       `json:"xp_earned"`
}

func (m *MeService) ReviewCard(userID, flashcardID uint64, quality int) (*ReviewResult, error) {
	if quality < 0 || quality > 5 {
		return nil, NewError(http.StatusUnprocessableEntity, "validation", "Rating harus 0-5.")
	}
	interval, reps, ease, err := m.store.ReviewByCard(userID, flashcardID)
	if err != nil {
		return nil, mapStoreErr(err)
	}
	newEase, newInterval, newReps := SM2(ease, interval, reps, quality)
	next := m.now().AddDate(0, 0, newInterval)
	if err = m.store.SaveReview(userID, flashcardID, newEase, newInterval, newReps, next); err != nil {
		return nil, mapStoreErr(err)
	}
	out := &ReviewResult{NextReviewAt: next, IntervalDays: newInterval, Repetitions: newReps, EaseFactor: newEase}
	if quality >= 3 {
		out.XPEarned = 2
		if _, err = m.store.Award(userID, 2, m.now()); err != nil {
			return nil, err
		}
	}
	return out, nil
}

func (m *MeService) Notes(userID, lessonID uint64) ([]model.Note, error) {
	return m.store.Notes(userID, lessonID)
}

func (m *MeService) UpsertNote(userID, lessonID uint64, content string) (*model.Note, error) {
	return m.store.UpsertNote(userID, lessonID, content)
}

func (m *MeService) UpdateNote(userID, noteID uint64, content string) error {
	return mapStoreErr(m.store.UpdateNote(userID, noteID, content))
}

func (m *MeService) DeleteNote(userID, noteID uint64) error {
	return mapStoreErr(m.store.DeleteNote(userID, noteID))
}

func (m *MeService) Bookmarks(userID uint64) ([]model.Bookmark, error) {
	return m.store.Bookmarks(userID)
}

func (m *MeService) AddBookmark(userID, lessonID uint64) error {
	return m.store.AddBookmark(userID, lessonID)
}

func (m *MeService) RemoveBookmark(userID, lessonID uint64) error {
	return m.store.RemoveBookmark(userID, lessonID)
}

func (m *MeService) Projects(userID uint64) ([]model.Project, error) {
	return m.store.Projects()
}

func (m *MeService) Project(slug string, userID uint64) (*model.Project, error) {
	p, err := m.store.ProjectBySlug(slug, userID)
	return p, mapStoreErr(err)
}

func (m *MeService) CompleteProjectStep(userID, stepID uint64) error {
	return m.store.CompleteProjectStep(userID, stepID, m.now())
}
