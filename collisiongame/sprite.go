package main

import "math/rand"

// Sprite and its friends
type Sprite interface {
	String() string
}

func RandNewSprite() Sprite {
	return [3]func() Sprite{
		func() Sprite { return NewHero() },
		func() Sprite { return NewWater() },
		func() Sprite { return NewFire() },
	}[rand.Intn(3)]()
}
