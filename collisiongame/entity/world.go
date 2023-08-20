package entity

import (
	"math/rand"
)

// World the happy sprites world
type World struct {
	Coord   [30]Sprite
	Handler IHandler
}

func NewWorld(h IHandler) *World {
	w := &World{Handler: h}
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
		w.Coord[num] = RandNewSprite()
	}
}

func (w *World) Move(x1 int, x2 int) {
	// TODO: isValidMove(x1, x2)

	c1Ptr := &w.Coord[x1]
	c2Ptr := &w.Coord[x2]

	// toCollide and move
	w.Handler.Handle(c1Ptr, c2Ptr)
}
