package store

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"golearn/internal/model"
)

func (s *Store) Paths() ([]model.Path, error) {
	var out []model.Path
	err := s.db.Select(&out, `SELECT p.*, COUNT(l.id) AS lesson_count
		FROM paths p
		LEFT JOIN modules m ON m.path_id = p.id
		LEFT JOIN lessons l ON l.module_id = m.id AND l.is_published = 1
		GROUP BY p.id ORDER BY p.order_no, p.id`)
	return out, err
}

func (s *Store) PathBySlug(slug string) (*model.Path, error) {
	var p model.Path
	err := s.db.Get(&p, `SELECT * FROM paths WHERE slug = ?`, slug)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return &p, err
}

func (s *Store) Modules(pathID uint64) ([]model.Module, error) {
	var out []model.Module
	err := s.db.Select(&out, `SELECT * FROM modules WHERE path_id = ? ORDER BY order_no, id`, pathID)
	return out, err
}

func (s *Store) LessonsByModule(moduleID uint64) ([]model.Lesson, error) {
	var out []model.Lesson
	err := s.db.Select(&out, `SELECT * FROM lessons WHERE module_id = ? AND is_published = 1 ORDER BY order_no, id`, moduleID)
	return out, err
}

func (s *Store) LessonBySlug(slug string) (*model.Lesson, *model.Module, *model.Path, error) {
	var l model.Lesson
	err := s.db.Get(&l, `SELECT * FROM lessons WHERE slug = ? AND is_published = 1`, slug)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil, nil, ErrNotFound
	}
	if err != nil {
		return nil, nil, nil, err
	}
	var m model.Module
	if err = s.db.Get(&m, `SELECT * FROM modules WHERE id = ?`, l.ModuleID); err != nil {
		return nil, nil, nil, err
	}
	var p model.Path
	if err = s.db.Get(&p, `SELECT * FROM paths WHERE id = ?`, m.PathID); err != nil {
		return nil, nil, nil, err
	}
	return &l, &m, &p, nil
}

// PathLessonList returns published lessons in learning order (module order, then lesson order).
func (s *Store) PathLessonList(pathID uint64) ([]model.Lesson, error) {
	var out []model.Lesson
	err := s.db.Select(&out, `SELECT l.* FROM lessons l
		JOIN modules m ON m.id = l.module_id
		WHERE m.path_id = ? AND l.is_published = 1
		ORDER BY m.order_no, m.id, l.order_no, l.id`, pathID)
	return out, err
}

func (s *Store) ExercisesByLesson(lessonID uint64) ([]model.Exercise, error) {
	var out []model.Exercise
	err := s.db.Select(&out, `SELECT * FROM exercises WHERE lesson_id = ? ORDER BY id`, lessonID)
	return out, err
}

func (s *Store) ExerciseByID(id uint64) (*model.Exercise, error) {
	var e model.Exercise
	err := s.db.Get(&e, `SELECT * FROM exercises WHERE id = ?`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	return &e, err
}

func (s *Store) ExerciseTests(exerciseID uint64) ([]model.ExerciseTest, error) {
	var out []model.ExerciseTest
	err := s.db.Select(&out, `SELECT * FROM exercise_tests WHERE exercise_id = ? ORDER BY id`, exerciseID)
	return out, err
}

func (s *Store) HasPassedExercise(userID, exerciseID uint64) (bool, error) {
	var n int
	err := s.db.Get(&n, `SELECT COUNT(*) FROM user_exercise_submissions
		WHERE user_id = ? AND exercise_id = ? AND is_passed = 1`, userID, exerciseID)
	return n > 0, err
}

// PriorPassedSubmission reports a pass before the most recent submission.
func (s *Store) PriorPassedSubmission(userID, exerciseID uint64) (bool, error) {
	var n int
	err := s.db.Get(&n, `SELECT COUNT(*) FROM user_exercise_submissions
		WHERE user_id = ? AND exercise_id = ? AND is_passed = 1
		AND id < (SELECT MAX(id) FROM user_exercise_submissions WHERE user_id = ? AND exercise_id = ?)`,
		userID, exerciseID, userID, exerciseID)
	return n > 0, err
}

func (s *Store) CreateSubmission(userID, exerciseID uint64, code string, passed bool, output string) error {
	_, err := s.db.Exec(`INSERT INTO user_exercise_submissions (user_id, exercise_id, code, is_passed, output)
		VALUES (?, ?, ?, ?, ?)`, userID, exerciseID, code, passed, output)
	return err
}

func (s *Store) QuizzesByLesson(lessonID uint64) ([]model.Quiz, error) {
	var quizzes []model.Quiz
	if err := s.db.Select(&quizzes, `SELECT * FROM quizzes WHERE lesson_id = ? ORDER BY order_no, id`, lessonID); err != nil {
		return nil, err
	}
	if len(quizzes) == 0 {
		return quizzes, nil
	}
	ids := make([]uint64, len(quizzes))
	for i := range quizzes {
		ids[i] = quizzes[i].ID
	}
	query, args, err := sqlxIn(`SELECT * FROM quiz_options WHERE quiz_id IN (?) ORDER BY order_no, id`, ids)
	if err != nil {
		return nil, err
	}
	var opts []model.QuizOption
	if err = s.db.Select(&opts, query, args...); err != nil {
		return nil, err
	}
	byQuiz := map[uint64][]model.QuizOption{}
	for _, o := range opts {
		byQuiz[o.QuizID] = append(byQuiz[o.QuizID], o)
	}
	for i := range quizzes {
		quizzes[i].Options = byQuiz[quizzes[i].ID]
	}
	return quizzes, nil
}

func (s *Store) QuizByID(id uint64) (*model.Quiz, error) {
	var q model.Quiz
	err := s.db.Get(&q, `SELECT * FROM quizzes WHERE id = ?`, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	err = s.db.Select(&q.Options, `SELECT * FROM quiz_options WHERE quiz_id = ? ORDER BY order_no, id`, id)
	return &q, err
}

func (s *Store) HasCorrectAttempt(userID, quizID uint64) (bool, error) {
	var n int
	err := s.db.Get(&n, `SELECT COUNT(*) FROM user_quiz_attempts WHERE user_id = ? AND quiz_id = ? AND is_correct = 1`, userID, quizID)
	return n > 0, err
}

func (s *Store) CreateQuizAttempt(userID, quizID uint64, correct bool) error {
	_, err := s.db.Exec(`INSERT INTO user_quiz_attempts (user_id, quiz_id, is_correct) VALUES (?, ?, ?)`, userID, quizID, correct)
	return err
}

func (s *Store) Search(q string, limit int) ([]model.SearchHit, error) {
	like := "%" + q + "%"
	var out []model.SearchHit
	err := s.db.Select(&out, `SELECT l.slug, l.title, l.summary, p.title AS path_title,
			SUBSTRING(l.content_md, GREATEST(1, LOCATE(?, l.content_md) - 80), 240) AS snippet
		FROM lessons l
		JOIN modules m ON m.id = l.module_id
		JOIN paths p ON p.id = m.path_id
		WHERE l.is_published = 1 AND (l.title LIKE ? OR l.summary LIKE ? OR l.content_md LIKE ?)
		ORDER BY CASE WHEN l.title LIKE ? THEN 0 ELSE 1 END, l.id
		LIMIT ?`, q, like, like, like, like, limit)
	if err != nil {
		return nil, fmt.Errorf("search: %w", err)
	}
	for i := range out {
		out[i].Snippet = strings.TrimSpace(strings.ReplaceAll(out[i].Snippet, "\n", " "))
	}
	return out, nil
}

// sqlxIn expands a single (?) inside an IN clause.
func sqlxIn(query string, ids []uint64) (string, []any, error) {
	if len(ids) == 0 {
		return "", nil, errors.New("empty ids")
	}
	placeholders := strings.TrimSuffix(strings.Repeat("?,", len(ids)), ",")
	args := make([]any, len(ids))
	for i, id := range ids {
		args[i] = id
	}
	return strings.Replace(query, "(?)", "("+placeholders+")", 1), args, nil
}
