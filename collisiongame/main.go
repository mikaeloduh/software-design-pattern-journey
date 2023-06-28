package main

import (
	"math/rand"
)

type Sprite struct {
}

type World struct {
	coord [30]Sprite
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
		w.coord[num] = Sprite{}
	}
}

func main() {
	println("hello world")
}
