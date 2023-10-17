package entity

import (
	"testing"
)

func Test_RPG(t *testing.T) {
	t.Run("test new RPG", func(t *testing.T) {
		g := NewRPG()
		g.Run()
	})
}
