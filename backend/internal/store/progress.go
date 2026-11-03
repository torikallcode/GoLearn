package store

import (
	"database/sql"
	"errors"
	"time"

	"golearn/internal/model"
)

func (s *Store) MarkLessonStarted(userID, lessonID uint64) error {
	_, err := s.db.Exec(`INSERT INTO user_lesson_progress (user_id, lesson_id, status)
		VALUES (?, ?, 'in_progress')
		ON DUPLICATE KEY UPDATE status = IF(status = 'completed', status, 'in_progress')`, userID, lessonID)
	return err
}

func (s *Store) CompleteLesson(userID, lessonID uint64, now time.Time) (bool, error) {
	res, err := s.db.Exec(`INSERT INTO user_lesson_progress (user_id, lesson_id, status, completed_at)
		VALUES (?, ?, 'completed', ?)
		ON DUPLICATE KEY UPDATE
			completed_at = IF(status = 'completed', completed_at, VALUES(completed_at)),
			status = 'completed'`, userID, lessonID, now)
	if err != nil {
		return false, err
	}
	affected, _ := res.RowsAffected()
	// affected == 1 means inserted; == 2 means updated because status changed.
	return affected > 0, nil
}

func (s *Store) CompletedLessonIDs(userID uint64) (map[uint64]bool, error) {
	rows, err := s.db.Queryx(`SELECT lesson_id FROM user_lesson_progress WHERE user_id = ? AND status = 'completed'`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	out := map[uint64]bool{}
	for rows.Next() {
		var id uint64
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		out[id] = true
	}
	return out, rows.Err()
}

func (s *Store) ProgressByPath(userID uint64) ([]model.PathProgress, error) {
	var out []model.PathProgress
	err := s.db.Select(&out, `SELECT p.slug, p.title, p.level,
			COUNT(l.id) AS total,
			COALESCE(SUM(CASE WHEN pr.status = 'completed' THEN 1 ELSE 0 END), 0) AS completed
		FROM paths p
		LEFT JOIN modules m ON m.path_id = p.id
		LEFT JOIN lessons l ON l.module_id = m.id AND l.is_published = 1
		LEFT JOIN user_lesson_progress pr ON pr.lesson_id = l.id AND pr.user_id = ?
		GROUP BY p.id ORDER BY p.order_no`, userID)
	if err != nil {
		return nil, err
	}
	for i := range out {
		out[i].Percent = percent(out[i].Completed, out[i].Total)
	}
	return out, nil
}

func percent(done, total int) int {
	if total <= 0 {
		return 0
	}
	return done * 100 / total
}

func (s *Store) RecentProgress(userID uint64, limit int) ([]model.ProgressRow, error) {
	var out []model.ProgressRow
	err := s.db.Select(&out, `SELECT pr.lesson_id, l.slug, l.title, p.slug AS path_slug, p.title AS path_title,
			pr.status, pr.completed_at
		FROM user_lesson_progress pr
		JOIN lessons l ON l.id = pr.lesson_id
		JOIN modules m ON m.id = l.module_id
		JOIN paths p ON p.id = m.path_id
		WHERE pr.user_id = ? AND pr.completed_at IS NOT NULL
		ORDER BY pr.completed_at DESC LIMIT ?`, userID, limit)
	return out, err
}

func (s *Store) CompletedToday(userID uint64, now time.Time) (int, error) {
	var n int
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	err := s.db.Get(&n, `SELECT COUNT(*) FROM user_lesson_progress
		WHERE user_id = ? AND status = 'completed' AND completed_at >= ?`, userID, start)
	return n, err
}

func (s *Store) TotalCompleted(userID uint64) (int, error) {
	var n int
	err := s.db.Get(&n, `SELECT COUNT(*) FROM user_lesson_progress WHERE user_id = ? AND status = 'completed'`, userID)
	return n, err
}

func (s *Store) TotalPublishedLessons() (int, error) {
	var n int
	err := s.db.Get(&n, `SELECT COUNT(*) FROM lessons WHERE is_published = 1`)
	return n, err
}

func (s *Store) LessonProgress(userID, lessonID uint64) (string, error) {
	var status string
	err := s.db.Get(&status, `SELECT status FROM user_lesson_progress WHERE user_id = ? AND lesson_id = ?`, userID, lessonID)
	if errors.Is(err, sql.ErrNoRows) {
		return "", nil
	}
	return status, err
}

// ---- spaced repetition ----

// SeedReviews creates due review rows for a lesson's flashcards after completion.
func (s *Store) SeedReviews(userID, lessonID uint64, now time.Time) error {
	_, err := s.db.Exec(`INSERT IGNORE INTO user_flashcard_reviews (user_id, flashcard_id, next_review_at)
		SELECT ?, id, ? FROM flashcards WHERE lesson_id = ?`, userID, now, lessonID)
	return err
}

func (s *Store) DueCards(userID uint64, now time.Time, limit int) ([]model.DueCard, error) {
	var out []model.DueCard
	err := s.db.Select(&out, `SELECT r.flashcard_id, f.front, f.back, l.slug AS lesson_slug, l.title AS lesson_title,
			r.interval_days, r.repetitions, r.ease_factor, r.next_review_at
		FROM user_flashcard_reviews r
		JOIN flashcards f ON f.id = r.flashcard_id
		JOIN lessons l ON l.id = f.lesson_id
		WHERE r.user_id = ? AND r.next_review_at <= ?
		ORDER BY r.next_review_at LIMIT ?`, userID, now, limit)
	return out, err
}

func (s *Store) DueCardCount(userID uint64, now time.Time) (int, error) {
	var n int
	err := s.db.Get(&n, `SELECT COUNT(*) FROM user_flashcard_reviews WHERE user_id = ? AND next_review_at <= ?`, userID, now)
	return n, err
}

func (s *Store) ReviewByCard(userID, flashcardID uint64) (interval, reps int, ease float64, err error) {
	err = s.db.QueryRowx(`SELECT interval_days, repetitions, ease_factor FROM user_flashcard_reviews
		WHERE user_id = ? AND flashcard_id = ?`, userID, flashcardID).Scan(&interval, &reps, &ease)
	if errors.Is(err, sql.ErrNoRows) {
		err = ErrNotFound
	}
	return
}

func (s *Store) SaveReview(userID, flashcardID uint64, ease float64, interval, reps int, next time.Time) error {
	res, err := s.db.Exec(`UPDATE user_flashcard_reviews
		SET ease_factor = ?, interval_days = ?, repetitions = ?, next_review_at = ?
		WHERE user_id = ? AND flashcard_id = ?`, ease, interval, reps, next, userID, flashcardID)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return ErrNotFound
	}
	return nil
}

// ---- notes ----

func (s *Store) Notes(userID uint64, lessonID uint64) ([]model.Note, error) {
	var out []model.Note
	if lessonID > 0 {
		err := s.db.Select(&out, `SELECT * FROM notes WHERE user_id = ? AND lesson_id = ?`, userID, lessonID)
		return out, err
	}
	err := s.db.Select(&out, `SELECT * FROM notes WHERE user_id = ? ORDER BY updated_at DESC`, userID)
	return out, err
}

func (s *Store) UpsertNote(userID, lessonID uint64, content string) (*model.Note, error) {
	_, err := s.db.Exec(`INSERT INTO notes (user_id, lesson_id, content) VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE content = VALUES(content)`, userID, lessonID, content)
	if err != nil {
		return nil, err
	}
	var n model.Note
	err = s.db.Get(&n, `SELECT * FROM notes WHERE user_id = ? AND lesson_id = ?`, userID, lessonID)
	return &n, err
}

func (s *Store) UpdateNote(userID, noteID uint64, content string) error {
	res, err := s.db.Exec(`UPDATE notes SET content = ? WHERE id = ? AND user_id = ?`, content, noteID, userID)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return ErrNotFound
	}
	return nil
}

func (s *Store) DeleteNote(userID, noteID uint64) error {
	res, err := s.db.Exec(`DELETE FROM notes WHERE id = ? AND user_id = ?`, noteID, userID)
	if err != nil {
		return err
	}
	if n, _ := res.RowsAffected(); n == 0 {
		return ErrNotFound
	}
	return nil
}

// ---- bookmarks ----

func (s *Store) Bookmarks(userID uint64) ([]model.Bookmark, error) {
	var out []model.Bookmark
	err := s.db.Select(&out, `SELECT b.lesson_id, l.slug AS lesson_slug, l.title AS lesson_title, l.summary, b.created_at
		FROM bookmarks b JOIN lessons l ON l.id = b.lesson_id
		WHERE b.user_id = ? ORDER BY b.created_at DESC`, userID)
	return out, err
}

func (s *Store) IsBookmarked(userID, lessonID uint64) (bool, error) {
	var n int
	err := s.db.Get(&n, `SELECT COUNT(*) FROM bookmarks WHERE user_id = ? AND lesson_id = ?`, userID, lessonID)
	return n > 0, err
}

func (s *Store) AddBookmark(userID, lessonID uint64) error {
	_, err := s.db.Exec(`INSERT IGNORE INTO bookmarks (user_id, lesson_id) VALUES (?, ?)`, userID, lessonID)
	return err
}

func (s *Store) RemoveBookmark(userID, lessonID uint64) error {
	_, err := s.db.Exec(`DELETE FROM bookmarks WHERE user_id = ? AND lesson_id = ?`, userID, lessonID)
	return err
}

// ---- projects ----

func (s *Store) Projects() ([]model.Project, error) {
	var out []model.Project
	err := s.db.Select(&out, `SELECT * FROM projects ORDER BY id`)
	return out, err
}

func (s *Store) ProjectBySlug(slug string, userID uint64) (*model.Project, error) {
	var p model.Project
	err := s.db.Get(&p, `SELECT * FROM projects WHERE slug = ?`, slug)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	err = s.db.Select(&p.Steps, `SELECT ps.*, COALESCE(upp.is_done, 0) AS is_done, upp.completed_at
		FROM project_steps ps
		LEFT JOIN user_project_progress upp ON upp.step_id = ps.id AND upp.user_id = ?
		WHERE ps.project_id = ? ORDER BY ps.order_no`, userID, p.ID)
	return &p, err
}

func (s *Store) CompleteProjectStep(userID, stepID uint64, now time.Time) error {
	_, err := s.db.Exec(`INSERT INTO user_project_progress (user_id, step_id, is_done, completed_at)
		VALUES (?, ?, 1, ?)
		ON DUPLICATE KEY UPDATE is_done = 1, completed_at = VALUES(completed_at)`, userID, stepID, now)
	return err
}
