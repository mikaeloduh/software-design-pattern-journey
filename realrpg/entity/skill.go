package entity

type ISkill interface {
	SelectTarget([]IUnit)
	Do()
}

type BasicAttack struct {
	Damage  int
	targets []IUnit
}

func (a *BasicAttack) SelectTarget(units []IUnit) {
	a.targets = units
}

func (a *BasicAttack) Do() {
	for _, unit := range a.targets {
		unit.SetHp(unit.GetHp() - a.Damage)
	}
}

type WaterBall struct {
	Damage  int
	targets []IUnit
}

func (a *WaterBall) SelectTarget(units []IUnit) {
	a.targets = units
}

func (a *WaterBall) Do() {
	for _, unit := range a.targets {
		unit.SetHp(unit.GetHp() - a.Damage)
	}
}
