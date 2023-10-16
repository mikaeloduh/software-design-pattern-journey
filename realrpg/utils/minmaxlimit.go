package utils

func MinMaxLimit(min, max, n int) int {
	if n < min {
		n = min
	} else if n > max {
		n = max
	}
	return n
}
