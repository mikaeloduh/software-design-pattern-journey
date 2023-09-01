package utils

import "math/rand"

// RandNonRepeatInt
func RandNonRepeatInt(min, max, count int) []int {
	if max-min+1 < count {
		return nil
	}

	numbers := make([]int, max-min+1)
	for i := min; i <= max; i++ {
		numbers[i-min] = i
	}

	for i := len(numbers) - 1; i >= len(numbers)-count; i-- {
		j := rand.Intn(i + 1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers[len(numbers)-count:]
}

// RandBool
func RandBool() bool {
	return rand.Intn(2) == 1
}
