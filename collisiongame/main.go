package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

// Sprite and its friends
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

func NewWater() *Water {
	return &Water{}
}

func (w *Water) String() string {
	return "W"
}

type Fire struct{}

func NewFire() *Fire {
	return &Fire{}
}

func (f *Fire) String() string {
	return "F"
}

// TypeSet is set of type
type TypeSet map[reflect.Type]bool

func NewTypeSet(c1, c2 any) TypeSet {
	return TypeSet{reflect.TypeOf(c1): true, reflect.TypeOf(c2): true}
}

func isSameTypeSet(s1, s2 TypeSet) bool {
	return reflect.DeepEqual(s1, s2)
}

// World the sprites' world
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
		w.coord[num] = randSprite()
	}
}

func randSprite() Sprite {
	return [3]func() Sprite{
		func() Sprite { return NewHero() },
		func() Sprite { return NewWater() },
		func() Sprite { return NewFire() },
	}[rand.Intn(3)]()
}

/**
 * | Subject \ target |  Hero  | Water  |  Fire  |
 * | :--------------: | :----: | :----: | :----: |
 * |       Hero       |  fail  |  +10   |  -10   |
 * |      Water       | remove |  fail  | remove |
 * |       Fire       | remove | remove |  fail  |
 */

func (w *World) Move(x1 int, x2 int) {
	//isValidMove(c1, c2)

	c1Ptr := &w.coord[x1]
	c2Ptr := &w.coord[x2]

	//toCollide(c1, c2)

	// toCollide and move
	if isSameType(*c1Ptr, &Hero{}) && isSameType(*c2Ptr, &Water{}) {
		// (Hero, Water)
		(*c1Ptr).(*Hero).SetHp(+10)
		*c2Ptr = nil
		*c2Ptr = *c1Ptr
		*c1Ptr = nil
	} else if isSameType(*c1Ptr, &Water{}) && isSameType(*c2Ptr, &Fire{}) {
		// (Water, Fire)
		*c1Ptr = nil
		*c2Ptr = nil
	} else if isSameType(*c1Ptr, &Fire{}) && isSameType(*c2Ptr, &Hero{}) {
		// (Fire, Hero)
		*c1Ptr = nil
		(*c2Ptr).(*Hero).SetHp(-10)
	} else if isSameType(*c2Ptr, *c2Ptr) {
		return
	}
}

func isSameType(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}

func main() {
	println("hello world")
}
