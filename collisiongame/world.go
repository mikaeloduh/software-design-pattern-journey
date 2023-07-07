package main

import "math/rand"

// World the happy sprites world
type World struct {
	coord   [30]Sprite
	handler IHandler
}

func NewWorld(h IHandler) *World {
	w := &World{handler: h}
	w.Init()
	return w
}

func (w *World) Init() {
	numbers := make([]int, 30)
	for i := 0; i < 30; i++ {
		numbers[i] = i
	}
	for i := len(numbers) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	for _, num := range numbers[:10] {
		w.coord[num] = RandNewSprite()
	}
}

func (w *World) Move(x1 int, x2 int) {
	// TODO: isValidMove(x1, x2)

	c1Ptr := &w.coord[x1]
	c2Ptr := &w.coord[x2]

	// toCollide and move
	w.handler.Handle(c1Ptr, c2Ptr)
}
