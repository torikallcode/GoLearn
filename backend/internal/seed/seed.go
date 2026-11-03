package seed

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"gopkg.in/yaml.v3"
)

//go:embed content/*.md content/*.yaml
var contentFS embed.FS

type lessonFile struct {
	Slug       string       `yaml:"slug"`
	Title      string       `yaml:"title"`
	Summary    string       `yaml:"summary"`
	Minutes    int          `yaml:"minutes"`
	Exercise   *exerciseDef `yaml:"exercise"`
	Quizzes    []quizDef    `yaml:"quizzes"`
	Flashcards []cardDef    `yaml:"flashcards"`
	Body       string       `yaml:"-"`
}

type exerciseDef struct {
	Title        string    `yaml:"title"`
	Difficulty   string    `yaml:"difficulty"`
	XP           int       `yaml:"xp"`
	Instructions string    `yaml:"instructions"`
	Starter      string    `yaml:"starter"`
	Solution     string    `yaml:"solution"`
	Hints        []string  `yaml:"hints"`
	Tests        []testDef `yaml:"tests"`
}

type testDef struct {
	Input    string `yaml:"input"`
	Expected string `yaml:"expected"`
	Hidden   bool   `yaml:"hidden"`
}

type quizDef struct {
	Question      string      `yaml:"question"`
	Type          string      `yaml:"type"`
	Explanation   string      `yaml:"explanation"`
	CorrectAnswer string      `yaml:"correct_answer"`
	Options       []optionDef `yaml:"options"`
}

type optionDef struct {
	Text    string `yaml:"text"`
	Correct bool   `yaml:"correct"`
}

type cardDef struct {
	Front string `yaml:"front"`
	Back  string `yaml:"back"`
}

type projectFile struct {
	Projects []projectDef `yaml:"projects"`
}

type projectDef struct {
	Slug        string        `yaml:"slug"`
	Title       string        `yaml:"title"`
	Description string        `yaml:"description"`
	Level       string        `yaml:"level"`
	Steps       []projectStep `yaml:"steps"`
}

type projectStep struct {
	Title   string `yaml:"title"`
	Content string `yaml:"content"`
}

func Run(db *sqlx.DB) error {
	lessons, err := loadLessons()
	if err != nil {
		return err
	}
	tx, err := db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, p := range foundationPaths(lessons) {
		if err = seedPath(tx, p); err != nil {
			return fmt.Errorf("seed path %s: %w", p.Slug, err)
		}
	}
	for _, p := range skeletonPaths() {
		if err = upsertSkeletonPath(tx, p); err != nil {
			return err
		}
	}
	if err = seedBadges(tx); err != nil {
		return err
	}
	if err = seedProjects(tx); err != nil {
		return fmt.Errorf("seed projects: %w", err)
	}
	return tx.Commit()
}

type seedPathDef struct {
	Slug, Title, Description, Level string
	Order                           int
	Modules                         []seedModuleDef
}

type seedModuleDef struct {
	Title   string
	Lessons []*lessonFile
}

func loadLessons() (map[string]*lessonFile, error) {
	entries, err := contentFS.ReadDir("content")
	if err != nil {
		return nil, err
	}
	out := map[string]*lessonFile{}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".md") {
			continue
		}
		raw, err := contentFS.ReadFile("content/" + e.Name())
		if err != nil {
			return nil, err
		}
		l, err := parseLesson(string(raw))
		if err != nil {
			return nil, fmt.Errorf("%s: %w", e.Name(), err)
		}
		out[l.Slug] = l
	}
	return out, nil
}

func parseLesson(raw string) (*lessonFile, error) {
	raw = strings.TrimSpace(raw)
	if !strings.HasPrefix(raw, "---") {
		return nil, errors.New("missing frontmatter")
	}
	parts := strings.SplitN(raw, "---", 3)
	if len(parts) != 3 {
		return nil, errors.New("invalid frontmatter")
	}
	var l lessonFile
	if err := yaml.Unmarshal([]byte(parts[1]), &l); err != nil {
		return nil, err
	}
	if l.Slug == "" || l.Title == "" {
		return nil, errors.New("slug and title are required")
	}
	if l.Minutes == 0 {
		l.Minutes = 10
	}
	l.Body = strings.TrimSpace(parts[2])
	return &l, nil
}

func seedPath(tx *sqlx.Tx, p seedPathDef) error {
	if _, err := tx.Exec(`INSERT INTO paths (slug, title, description, level, order_no)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE title=VALUES(title), description=VALUES(description), level=VALUES(level), order_no=VALUES(order_no)`,
		p.Slug, p.Title, p.Description, p.Level, p.Order); err != nil {
		return err
	}
	var pathID uint64
	if err := tx.Get(&pathID, `SELECT id FROM paths WHERE slug = ?`, p.Slug); err != nil {
		return err
	}
	lessonOrder := 0
	for mi, m := range p.Modules {
		if _, err := tx.Exec(`INSERT INTO modules (path_id, title, order_no) VALUES (?, ?, ?)
			ON DUPLICATE KEY UPDATE order_no=VALUES(order_no)`, pathID, m.Title, mi+1); err != nil {
			return err
		}
		var moduleID uint64
		if err := tx.Get(&moduleID, `SELECT id FROM modules WHERE path_id = ? AND title = ?`, pathID, m.Title); err != nil {
			return err
		}
		for _, l := range m.Lessons {
			lessonOrder++
			if err := seedLesson(tx, moduleID, lessonOrder, l); err != nil {
				return fmt.Errorf("lesson %s: %w", l.Slug, err)
			}
		}
	}
	return nil
}

func seedLesson(tx *sqlx.Tx, moduleID uint64, order int, l *lessonFile) error {
	if _, err := tx.Exec(`INSERT INTO lessons (module_id, slug, title, summary, content_md, estimated_minutes, order_no, is_published)
		VALUES (?, ?, ?, ?, ?, ?, ?, 1)
		ON DUPLICATE KEY UPDATE module_id=VALUES(module_id), title=VALUES(title), summary=VALUES(summary),
			content_md=VALUES(content_md), estimated_minutes=VALUES(estimated_minutes), order_no=VALUES(order_no), is_published=1`,
		moduleID, l.Slug, l.Title, l.Summary, l.Body, l.Minutes, order); err != nil {
		return err
	}
	var lessonID uint64
	if err := tx.Get(&lessonID, `SELECT id FROM lessons WHERE slug = ?`, l.Slug); err != nil {
		return err
	}
	if l.Exercise != nil {
		if err := seedExercise(tx, lessonID, l.Exercise); err != nil {
			return err
		}
	}
	for i, q := range l.Quizzes {
		if err := seedQuiz(tx, lessonID, i+1, q); err != nil {
			return err
		}
	}
	for _, c := range l.Flashcards {
		if _, err := tx.Exec(`INSERT INTO flashcards (lesson_id, front, back) VALUES (?, ?, ?)
			ON DUPLICATE KEY UPDATE back=VALUES(back)`, lessonID, c.Front, c.Back); err != nil {
			return err
		}
	}
	return nil
}

// marshalHints encodes hints as JSON because exercises.hints is a JSON column.
func marshalHints(hints []string) (string, error) {
	if hints == nil {
		return "[]", nil
	}
	b, err := json.Marshal(hints)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func seedExercise(tx *sqlx.Tx, lessonID uint64, e *exerciseDef) error {
	hints, err := marshalHints(e.Hints)
	if err != nil {
		return err
	}
	if e.XP == 0 {
		e.XP = 20
	}
	if e.Difficulty == "" {
		e.Difficulty = "mudah"
	}
	if _, err = tx.Exec(`INSERT INTO exercises (lesson_id, title, instructions_md, starter_code, solution_code, hints, difficulty, xp_reward)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE instructions_md=VALUES(instructions_md), starter_code=VALUES(starter_code),
			solution_code=VALUES(solution_code), hints=VALUES(hints), difficulty=VALUES(difficulty), xp_reward=VALUES(xp_reward)`,
		lessonID, e.Title, e.Instructions, e.Starter, e.Solution, string(hints), e.Difficulty, e.XP); err != nil {
		return err
	}
	var exID uint64
	if err = tx.Get(&exID, `SELECT id FROM exercises WHERE lesson_id = ? AND title = ?`, lessonID, e.Title); err != nil {
		return err
	}
	if _, err = tx.Exec(`DELETE FROM exercise_tests WHERE exercise_id = ?`, exID); err != nil {
		return err
	}
	for _, t := range e.Tests {
		if _, err = tx.Exec(`INSERT INTO exercise_tests (exercise_id, input, expected_output, is_hidden) VALUES (?, ?, ?, ?)`,
			exID, t.Input, t.Expected, t.Hidden); err != nil {
			return err
		}
	}
	return nil
}

func seedQuiz(tx *sqlx.Tx, lessonID uint64, order int, q quizDef) error {
	if q.Type == "" {
		q.Type = "multiple_choice"
	}
	var correct *string
	if q.CorrectAnswer != "" {
		correct = &q.CorrectAnswer
	}
	if _, err := tx.Exec(`INSERT INTO quizzes (lesson_id, question, type, correct_answer, explanation, order_no)
		VALUES (?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE question=VALUES(question), type=VALUES(type), correct_answer=VALUES(correct_answer),
			explanation=VALUES(explanation)`, lessonID, q.Question, q.Type, correct, q.Explanation, order); err != nil {
		return err
	}
	var quizID uint64
	if err := tx.Get(&quizID, `SELECT id FROM quizzes WHERE lesson_id = ? AND order_no = ?`, lessonID, order); err != nil {
		return err
	}
	if _, err := tx.Exec(`DELETE FROM quiz_options WHERE quiz_id = ?`, quizID); err != nil {
		return err
	}
	for i, o := range q.Options {
		if _, err := tx.Exec(`INSERT INTO quiz_options (quiz_id, text, is_correct, order_no) VALUES (?, ?, ?, ?)`,
			quizID, o.Text, o.Correct, i+1); err != nil {
			return err
		}
	}
	return nil
}

func upsertSkeletonPath(tx *sqlx.Tx, p seedPathDef) error {
	_, err := tx.Exec(`INSERT INTO paths (slug, title, description, level, order_no)
		VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE title=VALUES(title), description=VALUES(description), level=VALUES(level), order_no=VALUES(order_no)`,
		p.Slug, p.Title, p.Description, p.Level, p.Order)
	return err
}

func seedBadges(tx *sqlx.Tx) error {
	badges := []struct{ code, name, desc string }{
		{"first-lesson", "Langkah Pertama", "Menyelesaikan pelajaran pertama."},
		{"streak-7", "Konsisten 7 Hari", "Belajar 7 hari berturut-turut."},
		{"streak-30", "Konsisten 30 Hari", "Belajar 30 hari berturut-turut."},
		{"foundation-graduate", "Lulus Fondasi Go", "Menyelesaikan seluruh Path Fondasi Go."},
	}
	for _, b := range badges {
		if _, err := tx.Exec(`INSERT INTO badges (code, name, description) VALUES (?, ?, ?)
			ON DUPLICATE KEY UPDATE name=VALUES(name), description=VALUES(description)`, b.code, b.name, b.desc); err != nil {
			return err
		}
	}
	return nil
}

func seedProjects(tx *sqlx.Tx) error {
	raw, err := contentFS.ReadFile("content/projects.yaml")
	if err != nil {
		return err
	}
	var pf projectFile
	if err = yaml.Unmarshal(raw, &pf); err != nil {
		return err
	}
	for _, p := range pf.Projects {
		if _, err := tx.Exec(`INSERT INTO projects (slug, title, description, level) VALUES (?, ?, ?, ?)
			ON DUPLICATE KEY UPDATE title=VALUES(title), description=VALUES(description), level=VALUES(level)`,
			p.Slug, p.Title, p.Description, p.Level); err != nil {
			return err
		}
		var projectID uint64
		if err := tx.Get(&projectID, `SELECT id FROM projects WHERE slug = ?`, p.Slug); err != nil {
			return err
		}
		for i, s := range p.Steps {
			if _, err := tx.Exec(`INSERT INTO project_steps (project_id, title, content_md, order_no) VALUES (?, ?, ?, ?)
				ON DUPLICATE KEY UPDATE title=VALUES(title), content_md=VALUES(content_md)`,
				projectID, s.Title, s.Content, i+1); err != nil {
				return err
			}
		}
	}
	return nil
}
