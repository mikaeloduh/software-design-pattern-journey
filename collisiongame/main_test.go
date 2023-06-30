package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	world := World{}

	assert.IsType(t, World{}, world)

	world.Init()

	assert.Equal(t, 30, len(world.coord))
}

func TestCollision(t *testing.T) {

	t.Run("test Hero:Water and Hero +10 and moved, Water be removed", func(t *testing.T) {
		w := World{
			coord: [30]Sprite{NewHero(), &Water{}},
		}

		w.Move(0, 1)

		fmt.Printf("%v \n", w.coord[0])
		fmt.Printf("%v \n", w.coord[1])

		assert.Equal(t, 30+10, w.coord[1].(*Hero).hp)
		assert.Equal(t, nil, w.coord[0])
	})

	t.Run("test Water:Fire, Water and Fire should be removed", func(t *testing.T) {
		w := World{
			coord: [30]Sprite{&Fire{}, &Water{}},
		}

		w.Move(0, 1)

		fmt.Printf("%v \n", w.coord[0])
		fmt.Printf("%v \n", w.coord[1])

		assert.Equal(t, nil, w.coord[1])
		assert.Equal(t, nil, w.coord[0])
	})

	t.Run("test Fire:Hero, Hero -10 hp and Fire should be removed", func(t *testing.T) {
		w := World{
			coord: [30]Sprite{NewHero(), &Fire{}},
		}

		w.Move(0, 1)

		fmt.Printf("%v \n", w.coord[0])
		fmt.Printf("%v \n", w.coord[1])

		assert.Equal(t, nil, w.coord[0])
		assert.Equal(t, 30-10, w.coord[1].(*Hero).hp)
	})
}
