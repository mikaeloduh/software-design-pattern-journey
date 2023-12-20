package entity

import "math/rand"

type DefaultAI struct {
	seed int64
}

func NewDefaultAI() *DefaultAI {
	return &DefaultAI{seed: 0}
}

func (a *DefaultAI) IncrSeed() {
	a.seed++
}

func (a *DefaultAI) RandAction(num int) int {
	rand.Seed(a.seed)
	return rand.Int()
}
