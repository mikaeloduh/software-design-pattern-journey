package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollisionGame(t *testing.T) {
	t.Parallel()

	handler := NewHeroHeroHandler(NewHeroWaterHandler(NewHeroFireHandler(NewWaterHeroHandler(NewWaterWaterHandler(NewWaterFireHandler(NewFireHeroHandler(NewFireWaterHandler(NewFireFireHandler(nil)))))))))

	t.Run("New a sprites world, coord should have 30 units", func(t *testing.T) {
		w := NewWorld(handler)

		assert.Equal(t, 30, len(w.coord))
	})

	t.Run("Hero -> Hero, move fail", func(t *testing.T) {
		hero1 := NewHero()
		hero2 := NewHero()
		w := World{
			coord:   [30]Sprite{hero1, hero2},
			handler: handler,
		}

		w.Move(0, 1)

		assert.Same(t, hero1, w.coord[0])
		assert.Same(t, hero2, w.coord[1])
	})

	t.Run("Hero -> Water, Hero HP +10 and moved, Water should be removed", func(t *testing.T) {
		w := World{
			coord:   [30]Sprite{NewHero(), &Water{}},
			handler: handler,
		}

		w.Move(0, 1)

		fmt.Printf("%v \n", w.coord[0])
		fmt.Printf("%v \n", w.coord[1])
		fmt.Printf("%#v \n", w)

		assert.Equal(t, 30+10, w.coord[1].(*Hero).hp)
		assert.Equal(t, nil, w.coord[0])
	})

	t.Run("Hero -> Fire, Hero HP -10 and moved, Fire removed", func(t *testing.T) {
		hero := NewHero()
		fire := NewFire()
		w := World{
			coord:   [30]Sprite{hero, fire},
			handler: handler,
		}

		w.Move(0, 1)

		assert.Equal(t, 30-10, hero.hp)
		assert.Equal(t, nil, w.coord[0])
		assert.Same(t, hero, w.coord[1])
	})

	t.Run("Water -> Hero, Water removed, Hero HP +10 and moved", func(t *testing.T) {
		water := NewWater()
		hero := NewHero()
		w := World{
			coord:   [30]Sprite{water, hero},
			handler: handler,
		}

		w.Move(0, 1)

		assert.Equal(t, nil, w.coord[0])
		assert.Same(t, hero, w.coord[1])
		assert.Equal(t, 30+10, hero.hp)
	})

	t.Run("Water -> Water, move fail", func(t *testing.T) {
		water1 := NewWater()
		water2 := NewWater()
		w := World{
			coord:   [30]Sprite{water1, water2},
			handler: handler,
		}

		w.Move(0, 1)

		assert.Same(t, water1, w.coord[0])
		assert.Same(t, water2, w.coord[1])
	})

	t.Run("Water -> Fire, Water and Fire should be removed", func(t *testing.T) {
		w := World{
			coord:   [30]Sprite{&Water{}, &Fire{}},
			handler: handler,
		}

		w.Move(0, 1)

		assert.Equal(t, nil, w.coord[1])
		assert.Equal(t, nil, w.coord[0])
	})

	t.Run("Fire -> Hero, Hero -10 hp and Fire should be removed", func(t *testing.T) {
		hero := NewHero()
		w := World{
			coord:   [30]Sprite{&Fire{}, hero},
			handler: handler,
		}

		w.Move(0, 1)

		assert.Equal(t, 30-10, hero.hp)
		assert.Equal(t, nil, w.coord[0])
	})

	t.Run("Fire -> Water, both removed", func(t *testing.T) {
		w := World{
			coord:   [30]Sprite{&Fire{}, &Water{}},
			handler: handler,
		}

		w.Move(0, 1)

		assert.Equal(t, nil, w.coord[0])
		assert.Equal(t, nil, w.coord[1])
	})

	t.Run("Fire -> Fire, move fail", func(t *testing.T) {
		fire1 := NewFire()
		fire2 := NewFire()
		w := World{
			coord:   [30]Sprite{fire1, fire2},
			handler: handler,
		}

		w.Move(0, 1)

		assert.Same(t, fire1, w.coord[0])
		assert.Same(t, fire2, w.coord[1])
	})
}
