package entity

import (
	"math/rand"
	"time"
)

type AIPlayerInput struct{}

func (i AIPlayerInput) InputString() string {
	//TODO implement me
	panic("implement me")
}

func (AIPlayerInput) InputNum(min int, max int) int {
	rand.Seed(time.Now().UnixNano())

	return min + rand.Intn(max-min+1)
}

func (AIPlayerInput) InputBool() bool {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) == 1
}
