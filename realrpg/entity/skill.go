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
	units []IUnit
}

func (a *BasicAttack) SelectTarget(units []IUnit) {
	a.units = units
}

func (a *BasicAttack) Do() {
	for _, unit := range a.units {
		unit.SetHp(unit.GetHp() - 10)
	}
}
