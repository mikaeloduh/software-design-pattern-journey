package entity

type ISkill interface {
	SelectTarget([]IUnit)
	Do()
	GetMPCost() int
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

func (a *BasicAttack) GetMPCost() int {
	return 0
}

type WaterBall struct {
	Damage  int
	MPCost  int
	targets []IUnit
}

func (a *WaterBall) GetMPCost() int {
	return a.MPCost
}

func (a *WaterBall) SelectTarget(units []IUnit) {
	a.targets = units
}

func (a *WaterBall) Do() {
	for _, unit := range a.targets {
		unit.SetHp(unit.GetHp() - a.Damage)
	}
}
