package service

import "testing"

func TestSM2FirstCorrect(t *testing.T) {
	ease, interval, reps := SM2(2.5, 0, 0, 5)
	if interval != 1 {
		t.Fatalf("first correct review should schedule 1 day, got %d", interval)
	}
	if reps != 1 {
		t.Fatalf("expected reps=1, got %d", reps)
	}
	if ease <= 2.5 {
		t.Fatalf("perfect recall should increase ease, got %.2f", ease)
	}
}

func TestSM2SecondCorrect(t *testing.T) {
	_, interval, reps := SM2(2.5, 1, 1, 4)
	if interval != 6 {
		t.Fatalf("second correct review should schedule 6 days, got %d", interval)
	}
	if reps != 2 {
		t.Fatalf("expected reps=2, got %d", reps)
	}
}

func TestSM2LapseResets(t *testing.T) {
	ease, interval, reps := SM2(2.5, 6, 3, 1)
	if interval != 1 {
		t.Fatalf("lapse should reset interval to 1, got %d", interval)
	}
	if reps != 0 {
		t.Fatalf("lapse should reset reps to 0, got %d", reps)
	}
	if ease < 1.3 {
		t.Fatalf("ease must not drop below 1.3, got %.2f", ease)
	}
}

func TestSM2ClampsQuality(t *testing.T) {
	easeHigh, _, _ := SM2(2.5, 0, 0, 99)
	easePerfect, _, _ := SM2(2.5, 0, 0, 5)
	if easeHigh != easePerfect {
		t.Fatalf("quality above 5 should clamp to 5: %.2f != %.2f", easeHigh, easePerfect)
	}
}
