package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

type Sprite interface {
	String() string
}

type Hero struct {
	Sprite
	hp int
}

func (h *Hero) String() string {
	return fmt.Sprintf("H - %d", h.hp)
}

func (h *Hero) SetHp(n int) {
	h.hp += n
}

func NewHero() *Hero {
	return &Hero{hp: 30}
}

type Water struct{}

func (w *Water) String() string {
	return "W"
}

type Fire struct{}

func (f *Fire) String() string {
	return "F"
}

type World struct {
	coord [30]Sprite
}

type TypeSet map[reflect.Type]bool

func NewTypeSet(c1, c2 any) TypeSet {
	return TypeSet{reflect.TypeOf(c1): true, reflect.TypeOf(c2): true}
}

func isSameTypeSet(s1, s2 TypeSet) bool {
	return reflect.DeepEqual(s1, s2)
}

func (w *World) Move(x1 int, x2 int) {
	HeroWater := NewTypeSet(new(Hero), new(Water))
	WaterFire := NewTypeSet(new(Water), new(Fire))
	FireHero := NewTypeSet(new(Fire), new(Hero))

	c1 := w.coord[x1]
	c2 := w.coord[x2]

	ss := NewTypeSet(c1, c2)

	// (Hero, Water)
	if isSameTypeSet(HeroWater, ss) {
		var handle func(c Sprite, x int)
		handle = func(c Sprite, x int) {
			switch t := c.(type) {
			case *Hero:
				t.SetHp(10)
			case *Water:
				w.coord[x] = nil
			}
		}

		handle(c1, x1)
		handle(c2, x2)
		w.coord[x1], w.coord[x2] = w.coord[x2], w.coord[x1]
	} else if isSameTypeSet(WaterFire, ss) {
		var handle func(c Sprite, x int)
		handle = func(c Sprite, x int) {
			switch c.(type) {
			case *Water:
				w.coord[x] = nil
			case *Fire:
				w.coord[x] = nil
			}
		}

		handle(c1, x1)
		handle(c2, x2)
		w.coord[x1], w.coord[x2] = w.coord[x2], w.coord[x1]
	} else if isSameTypeSet(FireHero, ss) {
		var handle func(c Sprite, x int)
		handle = func(c Sprite, x int) {
			switch t := c.(type) {
			case *Fire:
				w.coord[x] = nil
			case *Hero:
				t.SetHp(-10)
			}
		}

		handle(c1, x1)
		handle(c2, x2)
		w.coord[x1], w.coord[x2] = w.coord[x2], w.coord[x1]
	} else if reflect.TypeOf(c1) == reflect.TypeOf(c2) {
		return
	}
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
		w.coord[num] = NewHero()
	}
}

func main() {
	println("hello world")
}
