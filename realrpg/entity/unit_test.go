package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHero(t *testing.T) {
	t.Run("test Hero BasicAttack", func(t *testing.T) {
		u1 := NewHero([]ISkill{&BasicAttack{}, WaterBall{}})
		u2 := NewHero([]ISkill{&BasicAttack{}})

		action1 := u1.Skills[0]
		action1.SelectTarget([]IUnit{u2})
		action1.Do()

		assert.Equal(t, 100-10, u2.HP)
	})
}
