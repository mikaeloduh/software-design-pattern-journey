package entity

type ISkill interface {
	SelectTarget([]IUnit)
	Do()
}

type WaterBall struct{}

func (a WaterBall) SelectTarget(units []IUnit) {
	//TODO implement me
	panic("implement me")
}

func (a WaterBall) Do() {
	//TODO implement me
	panic("implement me")
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
