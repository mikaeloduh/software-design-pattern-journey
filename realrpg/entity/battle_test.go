package entity

import "testing"

func TestBattle(t *testing.T) {
	t.Skip()

	t.Run("test run a 3v3 battle", func(t *testing.T) {
		hero := NewHero("Hero")
		hero.AddSkill(NewWaterBall(hero))

		ally1 := NewNonHero(NewDefaultAI())
		ally1.AddSkill(NewSelfExplosion(ally1))

		ally2 := NewNonHero(NewDefaultAI())
		ally2.AddSkill(NewCheerUp(ally2))

		enemy1 := NewNonHero(NewDefaultAI())
		enemy1.AddSkill(FakeNewOnePunch(enemy1))

		enemy2 := NewNonHero(NewDefaultAI())
		enemy2.AddSkill(NewSummon(enemy2))

		enemy3 := NewNonHero(NewDefaultAI())
		enemy3.AddSkill(NewCurse(enemy3))

		battle := NewBattle(
			NewTroop(hero, ally1, ally2),
			NewTroop(enemy1, enemy2, enemy3),
		)

		battle.Init()
		battle.Start()
		battle.GameResult()
	})
}

func FakeNewOnePunch(unit IUnit) *OnePunch {
	return &OnePunch{
		MPCost: 180,
		unit:   unit,
		handler: Hp500Handler{
			PetrochemicalStateHandler{
				PoisonedStateHandler{
					CheerUpStateHandler{
						NormalStateHandler{nil}}}}},
	}
}
