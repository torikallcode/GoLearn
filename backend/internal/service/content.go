package service

import (
	"encoding/json"
	"errors"
	"strings"

	"golearn/internal/model"
	"golearn/internal/store"
)

type ContentService struct {
	store *store.Store
}

func NewContentService(s *store.Store) *ContentService { return &ContentService{store: s} }

type LessonCard struct {
	ID               uint64 `json:"id"`
	Slug             string `json:"slug"`
	Title            string `json:"title"`
	Summary          string `json:"summary"`
	EstimatedMinutes int    `json:"estimated_minutes"`
	OrderNo          int    `json:"order_no"`
	Status           string `json:"status"`
}

type ModuleView struct {
	ID      uint64       `json:"id"`
	Title   string       `json:"title"`
	OrderNo int          `json:"order_no"`
	Lessons []LessonCard `json:"lessons"`
}

type PathDetail struct {
	model.Path
	Modules   []ModuleView `json:"modules"`
	Total     int          `json:"total"`
	Completed int          `json:"completed"`
	Percent   int          `json:"percent"`
}

type LessonDetail struct {
	Lesson     *model.Lesson    `json:"lesson"`
	Module     *model.Module    `json:"module"`
	Path       *model.Path      `json:"path"`
	Prev       *model.LessonNav `json:"prev"`
	Next       *model.LessonNav `json:"next"`
	Completed  bool             `json:"completed"`
	Bookmarked bool             `json:"bookmarked"`
	Note       *string          `json:"note"`
}

func (cs *ContentService) Paths() ([]model.Path, error) { return cs.store.Paths() }

func (cs *ContentService) PathDetail(slug string, userID uint64) (*PathDetail, error) {
	p, err := cs.store.PathBySlug(slug)
	if err != nil {
		return nil, mapStoreErr(err)
	}
	modules, err := cs.store.Modules(p.ID)
	if err != nil {
		return nil, err
	}
	all, err := cs.store.PathLessonList(p.ID)
	if err != nil {
		return nil, err
	}
	completed := map[uint64]bool{}
	if userID > 0 {
		completed, err = cs.store.CompletedLessonIDs(userID)
		if err != nil {
			return nil, err
		}
	}
	statuses := lessonStatuses(all, completed)
	detail := &PathDetail{Path: *p, Total: len(all)}
	for _, m := range modules {
		mv := ModuleView{ID: m.ID, Title: m.Title, OrderNo: m.OrderNo, Lessons: []LessonCard{}}
		for _, l := range all {
			if l.ModuleID != m.ID {
				continue
			}
			mv.Lessons = append(mv.Lessons, LessonCard{
				ID: l.ID, Slug: l.Slug, Title: l.Title, Summary: l.Summary,
				EstimatedMinutes: l.EstimatedMinutes, OrderNo: l.OrderNo, Status: statuses[l.ID],
			})
			if statuses[l.ID] == "completed" {
				detail.Completed++
			}
		}
		detail.Modules = append(detail.Modules, mv)
	}
	if detail.Total > 0 {
		detail.Percent = detail.Completed * 100 / detail.Total
	}
	return detail, nil
}

func (cs *ContentService) Lesson(slug string, userID uint64) (*LessonDetail, error) {
	l, m, p, err := cs.store.LessonBySlug(slug)
	if err != nil {
		return nil, mapStoreErr(err)
	}
	all, err := cs.store.PathLessonList(p.ID)
	if err != nil {
		return nil, err
	}
	completed := map[uint64]bool{}
	if userID > 0 {
		completed, err = cs.store.CompletedLessonIDs(userID)
		if err != nil {
			return nil, err
		}
		if !completed[l.ID] && lessonStatuses(all, completed)[l.ID] == "locked" {
			return nil, ErrLocked
		}
	}
	detail := &LessonDetail{Lesson: l, Module: m, Path: p, Completed: completed[l.ID]}
	for i := range all {
		if all[i].ID == l.ID {
			if i > 0 {
				detail.Prev = &model.LessonNav{Slug: all[i-1].Slug, Title: all[i-1].Title}
			}
			if i+1 < len(all) {
				detail.Next = &model.LessonNav{Slug: all[i+1].Slug, Title: all[i+1].Title}
			}
			break
		}
	}
	if userID > 0 {
		if detail.Bookmarked, err = cs.store.IsBookmarked(userID, l.ID); err != nil {
			return nil, err
		}
		notes, err := cs.store.Notes(userID, l.ID)
		if err != nil {
			return nil, err
		}
		if len(notes) > 0 {
			detail.Note = &notes[0].Content
		}
	}
	return detail, nil
}

func (cs *ContentService) Exercises(slug string, userID uint64) ([]model.Exercise, error) {
	l, _, _, err := cs.store.LessonBySlug(slug)
	if err != nil {
		return nil, mapStoreErr(err)
	}
	out, err := cs.store.ExercisesByLesson(l.ID)
	if err != nil {
		return nil, err
	}
	for i := range out {
		setHints(&out[i])
		if userID > 0 {
			out[i].Passed, err = cs.store.HasPassedExercise(userID, out[i].ID)
			if err != nil {
				return nil, err
			}
		}
		// Reveal the reference solution only after the user has passed.
		if !out[i].Passed {
			out[i].SolutionCode = ""
		}
	}
	return out, nil
}

func (cs *ContentService) Quizzes(slug string) ([]model.Quiz, error) {
	l, _, _, err := cs.store.LessonBySlug(slug)
	if err != nil {
		return nil, mapStoreErr(err)
	}
	return cs.store.QuizzesByLesson(l.ID)
}

func (cs *ContentService) Search(q string) ([]model.SearchHit, error) {
	q = strings.TrimSpace(q)
	if len(q) < 2 {
		return []model.SearchHit{}, nil
	}
	return cs.store.Search(q, 20)
}

func lessonStatuses(lessons []model.Lesson, completed map[uint64]bool) map[uint64]string {
	out := make(map[uint64]string, len(lessons))
	blocked := false
	for _, l := range lessons {
		switch {
		case completed[l.ID]:
			out[l.ID] = "completed"
		case blocked:
			out[l.ID] = "locked"
		default:
			out[l.ID] = "available"
			blocked = true
		}
	}
	return out
}

func setHints(e *model.Exercise) {
	e.Hints = []string{}
	if strings.TrimSpace(e.HintsRaw) == "" {
		return
	}
	if err := json.Unmarshal([]byte(e.HintsRaw), &e.Hints); err != nil {
		e.Hints = []string{}
	}
}

func mapStoreErr(err error) error {
	if errors.Is(err, store.ErrNotFound) {
		return ErrNotFound
	}
	return err
}
