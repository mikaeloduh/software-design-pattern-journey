package entity

import (
	"testing"
)

func Test_RPG(t *testing.T) {
	t.Run("test new RPG", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit2 := NewHero("p2")
		g := NewRPG([]IUnit{unit1, unit2})
		g.Run()
	})
}
