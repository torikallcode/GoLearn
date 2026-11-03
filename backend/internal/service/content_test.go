package service

import (
	"testing"

	"golearn/internal/model"
)

func TestLessonStatusesSequentialUnlock(t *testing.T) {
	lessons := []model.Lesson{{ID: 1}, {ID: 2}, {ID: 3}}
	completed := map[uint64]bool{1: true}

	got := lessonStatuses(lessons, completed)

	if got[1] != "completed" {
		t.Fatalf("lesson 1 should be completed, got %q", got[1])
	}
	if got[2] != "available" {
		t.Fatalf("lesson 2 should be available, got %q", got[2])
	}
	if got[3] != "locked" {
		t.Fatalf("lesson 3 should be locked, got %q", got[3])
	}
}

func TestLessonStatusesAllCompleted(t *testing.T) {
	lessons := []model.Lesson{{ID: 1}, {ID: 2}}
	completed := map[uint64]bool{1: true, 2: true}

	got := lessonStatuses(lessons, completed)

	for _, l := range lessons {
		if got[l.ID] != "completed" {
			t.Fatalf("lesson %d should be completed, got %q", l.ID, got[l.ID])
		}
	}
}
