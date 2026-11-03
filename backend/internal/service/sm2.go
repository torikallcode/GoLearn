package service

// SM2 returns the next spaced-repetition state using the SM-2 algorithm.
// quality is 0..5 (0 = blank, 5 = perfect recall).
func SM2(ease float64, interval, reps, quality int) (float64, int, int) {
	if quality < 0 {
		quality = 0
	}
	if quality > 5 {
		quality = 5
	}
	if ease < 1.3 {
		ease = 2.5
	}
	if quality < 3 {
		reps = 0
		interval = 1
	} else {
		switch reps {
		case 0:
			interval = 1
		case 1:
			interval = 6
		default:
			interval = int(float64(interval)*ease + 0.5)
		}
		reps++
	}
	ease += 0.1 - float64(5-quality)*(0.08+float64(5-quality)*0.02)
	if ease < 1.3 {
		ease = 1.3
	}
	if interval < 1 {
		interval = 1
	}
	return ease, interval, reps
}
