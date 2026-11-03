package model

import "time"

type User struct {
	ID           uint64    `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Email        string    `db:"email" json:"email"`
	PasswordHash string    `db:"password_hash" json:"-"`
	Role         string    `db:"role" json:"role"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type Stats struct {
	UserID         uint64     `db:"user_id" json:"-"`
	XP             int        `db:"xp" json:"xp"`
	Level          int        `db:"level" json:"level"`
	CurrentStreak  int        `db:"current_streak" json:"current_streak"`
	LongestStreak  int        `db:"longest_streak" json:"longest_streak"`
	LastActiveDate *time.Time `db:"last_active_date" json:"last_active_date"`
}

type Path struct {
	ID          uint64 `db:"id" json:"id"`
	Slug        string `db:"slug" json:"slug"`
	Title       string `db:"title" json:"title"`
	Description string `db:"description" json:"description"`
	Level       string `db:"level" json:"level"`
	OrderNo     int    `db:"order_no" json:"order_no"`
	LessonCount int    `db:"lesson_count" json:"lesson_count"`
}

type Module struct {
	ID      uint64 `db:"id" json:"id"`
	PathID  uint64 `db:"path_id" json:"path_id"`
	Title   string `db:"title" json:"title"`
	OrderNo int    `db:"order_no" json:"order_no"`
}

type Lesson struct {
	ID               uint64 `db:"id" json:"id"`
	ModuleID         uint64 `db:"module_id" json:"module_id"`
	Slug             string `db:"slug" json:"slug"`
	Title            string `db:"title" json:"title"`
	Summary          string `db:"summary" json:"summary"`
	ContentMD        string `db:"content_md" json:"content_md,omitempty"`
	EstimatedMinutes int    `db:"estimated_minutes" json:"estimated_minutes"`
	OrderNo          int    `db:"order_no" json:"order_no"`
	IsPublished      bool   `db:"is_published" json:"is_published"`
}

type LessonNav struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
}

type Exercise struct {
	ID             uint64   `db:"id" json:"id"`
	LessonID       uint64   `db:"lesson_id" json:"lesson_id"`
	Title          string   `db:"title" json:"title"`
	InstructionsMD string   `db:"instructions_md" json:"instructions_md"`
	StarterCode    string   `db:"starter_code" json:"starter_code"`
	SolutionCode   string   `db:"solution_code" json:"solution_code,omitempty"`
	Hints          []string `db:"-" json:"hints"`
	HintsRaw       string   `db:"hints" json:"-"`
	Difficulty     string   `db:"difficulty" json:"difficulty"`
	XPReward       int      `db:"xp_reward" json:"xp_reward"`
	Passed         bool     `db:"passed" json:"passed"`
}

type ExerciseTest struct {
	ID             uint64 `db:"id" json:"id"`
	ExerciseID     uint64 `db:"exercise_id" json:"-"`
	Input          string `db:"input" json:"-"`
	ExpectedOutput string `db:"expected_output" json:"-"`
	IsHidden       bool   `db:"is_hidden" json:"is_hidden"`
}

type Quiz struct {
	ID            uint64       `db:"id" json:"id"`
	LessonID      uint64       `db:"lesson_id" json:"lesson_id"`
	Question      string       `db:"question" json:"question"`
	Type          string       `db:"type" json:"type"`
	CorrectAnswer *string      `db:"correct_answer" json:"-"`
	Explanation   string       `db:"explanation" json:"explanation"`
	OrderNo       int          `db:"order_no" json:"order_no"`
	Options       []QuizOption `db:"-" json:"options"`
}

type QuizOption struct {
	ID      uint64 `db:"id" json:"id"`
	QuizID  uint64 `db:"quiz_id" json:"-"`
	Text    string `db:"text" json:"text"`
	Correct bool   `db:"is_correct" json:"-"`
	OrderNo int    `db:"order_no" json:"-"`
}

type Flashcard struct {
	ID       uint64 `db:"id" json:"id"`
	LessonID uint64 `db:"lesson_id" json:"lesson_id"`
	Front    string `db:"front" json:"front"`
	Back     string `db:"back" json:"back"`
}

type DueCard struct {
	FlashcardID  uint64    `db:"flashcard_id" json:"flashcard_id"`
	Front        string    `db:"front" json:"front"`
	Back         string    `db:"back" json:"back"`
	LessonSlug   string    `db:"lesson_slug" json:"lesson_slug"`
	LessonTitle  string    `db:"lesson_title" json:"lesson_title"`
	IntervalDays int       `db:"interval_days" json:"interval_days"`
	Repetitions  int       `db:"repetitions" json:"repetitions"`
	EaseFactor   float64   `db:"ease_factor" json:"ease_factor"`
	NextReviewAt time.Time `db:"next_review_at" json:"next_review_at"`
}

type Note struct {
	ID        uint64     `db:"id" json:"id"`
	UserID    uint64     `db:"user_id" json:"-"`
	LessonID  uint64     `db:"lesson_id" json:"lesson_id"`
	Content   string     `db:"content" json:"content"`
	UpdatedAt time.Time  `db:"updated_at" json:"updated_at"`
	Lesson    *LessonNav `db:"-" json:"lesson,omitempty"`
}

type Bookmark struct {
	LessonID    uint64    `db:"lesson_id" json:"lesson_id"`
	LessonSlug  string    `db:"lesson_slug" json:"lesson_slug"`
	LessonTitle string    `db:"lesson_title" json:"lesson_title"`
	Summary     string    `db:"summary" json:"summary"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type Project struct {
	ID          uint64        `db:"id" json:"id"`
	Slug        string        `db:"slug" json:"slug"`
	Title       string        `db:"title" json:"title"`
	Description string        `db:"description" json:"description"`
	Level       string        `db:"level" json:"level"`
	Steps       []ProjectStep `db:"-" json:"steps,omitempty"`
}

type ProjectStep struct {
	ID          uint64     `db:"id" json:"id"`
	ProjectID   uint64     `db:"project_id" json:"-"`
	Title       string     `db:"title" json:"title"`
	ContentMD   string     `db:"content_md" json:"content_md"`
	OrderNo     int        `db:"order_no" json:"order_no"`
	IsDone      bool       `db:"is_done" json:"is_done"`
	CompletedAt *time.Time `db:"completed_at" json:"completed_at"`
}

type ProgressRow struct {
	LessonID    uint64     `db:"lesson_id" json:"lesson_id"`
	Slug        string     `db:"slug" json:"slug"`
	Title       string     `db:"title" json:"title"`
	PathSlug    string     `db:"path_slug" json:"path_slug"`
	PathTitle   string     `db:"path_title" json:"path_title"`
	Status      string     `db:"status" json:"status"`
	CompletedAt *time.Time `db:"completed_at" json:"completed_at"`
}

type Dashboard struct {
	Stats          Stats           `json:"stats"`
	DueReviews     int             `json:"due_reviews"`
	CompletedToday int             `json:"completed_today"`
	TotalCompleted int             `json:"total_completed"`
	TotalLessons   int             `json:"total_lessons"`
	Continue       *ContinueLesson `json:"continue"`
	Recent         []ProgressRow   `json:"recent"`
	Paths          []PathProgress  `json:"paths"`
}

type ContinueLesson struct {
	Slug     string `json:"slug"`
	Title    string `json:"title"`
	PathSlug string `json:"path_slug"`
	Minutes  int    `json:"minutes"`
}

type PathProgress struct {
	Slug      string `db:"slug" json:"slug"`
	Title     string `db:"title" json:"title"`
	Level     string `db:"level" json:"level"`
	Total     int    `db:"total" json:"total"`
	Completed int    `db:"completed" json:"completed"`
	Percent   int    `json:"percent"`
}

type SearchHit struct {
	Slug      string `db:"slug" json:"slug"`
	Title     string `db:"title" json:"title"`
	Summary   string `db:"summary" json:"summary"`
	PathTitle string `db:"path_title" json:"path_title"`
	Snippet   string `db:"snippet" json:"snippet"`
}
