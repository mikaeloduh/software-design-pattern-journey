package utils

import "math/rand"

func RandNonRepeatIntStack(min, max, count int) Stack[int] {
	if max-min+1 < count {
		return nil
	}

	numbers := NewStack[int]()
	for i := min; i <= max; i++ {
		numbers.Push(i)
	}

	for i := len(*numbers) - 1; i >= len(*numbers)-count; i-- {
		j := rand.Intn(i + 1)
		(*numbers)[i], (*numbers)[j] = (*numbers)[j], (*numbers)[i]
	}

	return (*numbers)[len(*numbers)-count:]
}

// RandBool
func RandBool() bool {
	return rand.Intn(2) == 1
}
