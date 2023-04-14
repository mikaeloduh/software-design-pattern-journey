package entity

import (
	"math/rand"
	"time"
)

type Input interface {
	InputNum(int, int) int
	InputBool() bool
}

type AIInput struct{}

func (AIInput) InputNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return min + rand.Intn(max-min+1)
}

func (AIInput) InputBool() bool {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) == 1
}

type UserInput struct{}

func (i UserInput) InputNum(min int, max int) int {
	// TODO: stdin
	rand.Seed(time.Now().UnixNano())

	return min + rand.Intn(max-min+1)
}

func (i UserInput) InputBool() bool {
	// TODO: stdin
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) == 1
}
