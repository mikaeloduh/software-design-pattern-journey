package utils

import (
	"errors"
	"math/rand"
)

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}

	index := len(*s) - 1
	item := (*s)[index]
	*s = (*s)[:index]

	return item, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

type NonRepeatInt Stack[int]

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
