package mathlib

// Sign function returns the sign (+1, -1 or 0) of a supplied number. ie.
// if the number is positive, the Sign function returns +1, if the
// number is negative, the function returns -1 and if the number is 0 (zero), the function returns 0.
func Sign(number float64) float64 {
	switch {
	case number < 0:
		return -1
	case number == 0:
		return 0
	}
	return 1
}
