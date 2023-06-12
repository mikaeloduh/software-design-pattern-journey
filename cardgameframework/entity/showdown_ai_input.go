package entity

import (
	"math/rand"
	"time"
)

type ShowdownAIInput struct{}

func (i ShowdownAIInput) InputString() string {
	//TODO implement me
	panic("implement me")
}

func (ShowdownAIInput) InputNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return min + rand.Intn(max-min+1)
}

func (ShowdownAIInput) InputBool() bool {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) == 1
}
