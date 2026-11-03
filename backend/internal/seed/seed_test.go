package seed

import (
	"encoding/json"
	"testing"
)

func TestContentParses(t *testing.T) {
	lessons, err := loadLessons()
	if err != nil {
		t.Fatalf("loadLessons: %v", err)
	}
	if len(lessons) < 11 {
		t.Fatalf("expected at least 11 lessons, got %d", len(lessons))
	}
	if _, err := contentFS.ReadFile("content/projects.yaml"); err != nil {
		t.Fatalf("read projects.yaml: %v", err)
	}
	for slug, l := range lessons {
		if l.Title == "" {
			t.Errorf("%s: missing title", slug)
		}
		if l.Body == "" {
			t.Errorf("%s: empty body", slug)
		}
		if l.Exercise == nil {
			t.Errorf("%s: missing exercise", slug)
		}
		if len(l.Quizzes) == 0 {
			t.Errorf("%s: missing quizzes", slug)
		}
		if len(l.Flashcards) == 0 {
			t.Errorf("%s: missing flashcards", slug)
		}
	}
}

func TestMarshalHintsProducesValidJSON(t *testing.T) {
	cases := [][]string{nil, {}, {"satu", "dua"}}
	for _, hints := range cases {
		out, err := marshalHints(hints)
		if err != nil {
			t.Fatalf("marshalHints: %v", err)
		}
		if !json.Valid([]byte(out)) {
			t.Fatalf("hints must be valid JSON for the JSON column, got %q", out)
		}
	}
}

func TestFoundationPathHasNoDanglingLessons(t *testing.T) {
	lessons, err := loadLessons()
	if err != nil {
		t.Fatalf("loadLessons: %v", err)
	}
	paths := foundationPaths(lessons)
	if len(paths) != 1 {
		t.Fatalf("expected one foundation path, got %d", len(paths))
	}
	total := 0
	for _, m := range paths[0].Modules {
		if len(m.Lessons) == 0 {
			t.Errorf("module %q has no lessons", m.Title)
		}
		total += len(m.Lessons)
	}
	if total != len(lessons) {
		t.Fatalf("foundation path references %d lessons but %d parsed", total, len(lessons))
	}
}
