package entity

import (
	"math/rand"
	"time"
)

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
