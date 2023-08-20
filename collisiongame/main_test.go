package main

import (
	"collisiongame/entity"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollisionGame(t *testing.T) {
	t.Parallel()

	handler := entity.NewHeroHeroHandler(
		entity.NewHeroWaterHandler(
			entity.NewHeroFireHandler(
				entity.NewWaterHeroHandler(
					entity.NewWaterWaterHandler(
						entity.NewWaterFireHandler(
							entity.NewFireHeroHandler(
								entity.NewFireWaterHandler(
									entity.NewFireFireHandler(nil)))))))))

	t.Run("New a sprites world, coord should have 30 units", func(t *testing.T) {
		w := entity.NewWorld(handler)

		assert.Equal(t, 30, len(w.Coord))
	})

	t.Run("Hero -> Hero, move fail", func(t *testing.T) {
		hero1 := entity.NewHero()
		hero2 := entity.NewHero()
		w := entity.World{
			Coord:   [30]entity.Sprite{hero1, hero2},
			Handler: handler,
		}

		w.Move(0, 1)

		assert.Same(t, hero1, w.Coord[0])
		assert.Same(t, hero2, w.Coord[1])
	})

	t.Run("Hero -> Water, Hero HP +10 and moved, Water should be removed", func(t *testing.T) {
		w := entity.World{
			Coord:   [30]entity.Sprite{entity.NewHero(), &entity.Water{}},
			Handler: handler,
		}

		w.Move(0, 1)

		fmt.Printf("%v \n", w.Coord[0])
		fmt.Printf("%v \n", w.Coord[1])
		fmt.Printf("%#v \n", w)

		assert.Equal(t, 30+10, w.Coord[1].(*entity.Hero).Hp)
		assert.Equal(t, nil, w.Coord[0])
	})

	t.Run("Hero -> Fire, Hero HP -10 and moved, Fire removed", func(t *testing.T) {
		hero := entity.NewHero()
		fire := entity.NewFire()
		w := entity.World{
			Coord:   [30]entity.Sprite{hero, fire},
			Handler: handler,
		}

		w.Move(0, 1)

		assert.Equal(t, 30-10, hero.Hp)
		assert.Equal(t, nil, w.Coord[0])
		assert.Same(t, hero, w.Coord[1])
	})

	t.Run("Water -> Hero, Water removed, Hero HP +10 and moved", func(t *testing.T) {
		water := entity.NewWater()
		hero := entity.NewHero()
		w := entity.World{
			Coord:   [30]entity.Sprite{water, hero},
			Handler: handler,
		}

		w.Move(0, 1)

		assert.Equal(t, nil, w.Coord[0])
		assert.Same(t, hero, w.Coord[1])
		assert.Equal(t, 30+10, hero.Hp)
	})

	t.Run("Water -> Water, move fail", func(t *testing.T) {
		water1 := entity.NewWater()
		water2 := entity.NewWater()
		w := entity.World{
			Coord:   [30]entity.Sprite{water1, water2},
			Handler: handler,
		}

		w.Move(0, 1)

		assert.Same(t, water1, w.Coord[0])
		assert.Same(t, water2, w.Coord[1])
	})

	t.Run("Water -> Fire, Water and Fire should be removed", func(t *testing.T) {
		w := entity.World{
			Coord:   [30]entity.Sprite{&entity.Water{}, &entity.Fire{}},
			Handler: handler,
		}

		w.Move(0, 1)

		assert.Equal(t, nil, w.Coord[1])
		assert.Equal(t, nil, w.Coord[0])
	})

	t.Run("Fire -> Hero, Hero -10 hp and Fire should be removed", func(t *testing.T) {
		hero := entity.NewHero()
		w := entity.World{
			Coord:   [30]entity.Sprite{&entity.Fire{}, hero},
			Handler: handler,
		}

		w.Move(0, 1)

		assert.Equal(t, 30-10, hero.Hp)
		assert.Equal(t, nil, w.Coord[0])
	})

	t.Run("Fire -> Water, both removed", func(t *testing.T) {
		w := entity.World{
			Coord:   [30]entity.Sprite{&entity.Fire{}, &entity.Water{}},
			Handler: handler,
		}

		w.Move(0, 1)

		assert.Equal(t, nil, w.Coord[0])
		assert.Equal(t, nil, w.Coord[1])
	})

	t.Run("Fire -> Fire, move fail", func(t *testing.T) {
		fire1 := entity.NewFire()
		fire2 := entity.NewFire()
		w := entity.World{
			Coord:   [30]entity.Sprite{fire1, fire2},
			Handler: handler,
		}

		w.Move(0, 1)

		assert.Same(t, fire1, w.Coord[0])
		assert.Same(t, fire2, w.Coord[1])
	})
}
