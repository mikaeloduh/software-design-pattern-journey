package entity

import (
	"math/rand"
	"time"
)

type AIInput struct{}

func (i AIInput) InputString() string {
	//TODO implement me
	panic("implement me")
}

func (AIInput) InputNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return min + rand.Intn(max-min+1)
}

func (AIInput) InputBool() bool {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) == 1
}
