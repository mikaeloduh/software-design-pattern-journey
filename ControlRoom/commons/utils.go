package commons

import (
	"errors"
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

func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		return *new(T), errors.New("stack is empty")
	}

	return (*s)[len(*s)-1], nil
}

func (s *Stack[T]) Clear() {
	*s = Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}
