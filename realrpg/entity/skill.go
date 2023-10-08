package entity

type ISkill interface {
	SelectTarget([]IUnit)
	Do()
	GetMPCost() int
	IsMpEnough() bool
}

// BasicAttack
type BasicAttack struct {
	unit    IUnit
	targets []IUnit
}

func NewBasicAttack(unit IUnit) *BasicAttack {
	return &BasicAttack{unit: unit}
}

func (a *BasicAttack) IsMpEnough() bool {
	return true
}

func (a *BasicAttack) SelectTarget(units []IUnit) {
	a.targets = units
}

func (a *BasicAttack) Do() {
	damage := a.unit.GetState().OnAttack(a.unit.GetSTR())
	for _, target := range a.targets {
		target.TakeDamage(damage)
	}
}

func (a *BasicAttack) GetMPCost() int {
	return 0
}

// WaterBall
type WaterBall struct {
	Damage  int
	MPCost  int
	unit    IUnit
	targets []IUnit
}

func NewWaterBall(unit IUnit) *WaterBall {
	return &WaterBall{
		Damage: 120,
		MPCost: 50,
		unit:   unit,
	}
}

func (a *WaterBall) IsMpEnough() bool {
	if a.unit.GetMp() < a.MPCost {
		return false
	}
	return true
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

	a.unit.ConsumeMp(a.MPCost)
}

//
//type Summon struct {
//	MPCost  int
//	targets []IUnit
//}
//
//func (a *Summon) SelectTarget(_ []IUnit) {
//	panic("invalid operation")
//}
//
//func (a *Summon) Do() {
//	slime := NewSlime()
//
//}
//
//func (a *Summon) GetMPCost() int {
//	//TODO implement me
//	panic("implement me")
//}

// SelfExplosion
type SelfExplosion struct {
	MPCost  int
	Damage  int
	unit    IUnit
	targets []IUnit
}

func (a *SelfExplosion) IsMpEnough() bool {
	if a.unit.GetMp() < a.MPCost {
		return false
	}
	return true
}

func NewSelfExplosion(unit IUnit) *SelfExplosion {
	return &SelfExplosion{
		MPCost: 200,
		Damage: 150,
		unit:   unit,
	}
}

func (a *SelfExplosion) SelectTarget(unit []IUnit) {
	a.targets = unit
}

func (a *SelfExplosion) Do() {
	for _, target := range a.targets {
		target.TakeDamage(a.Damage)
	}

	a.unit.SetHp(0)
}

func (a *SelfExplosion) GetMPCost() int {
	return a.MPCost
}

// CheerUp
type CheerUp struct {
	Damage  int
	MPCost  int
	unit    IUnit
	targets []IUnit
}

func (a *CheerUp) IsMpEnough() bool {
	if a.unit.GetMp() < a.MPCost {
		return false
	}
	return true
}

func (a *CheerUp) GetMPCost() int {
	return a.MPCost
}

func (a *CheerUp) SelectTarget(units []IUnit) {
	a.targets = units
}

func (a *CheerUp) Do() {
	for _, target := range a.targets {
		target.SetState(NewCheerUpState(target))
	}
}

// SelfHealing
type SelfHealing struct {
	Damage int
	MPCost int
	unit   IUnit
}

func NewSelfHealing(unit IUnit) *SelfHealing {
	return &SelfHealing{
		Damage: -50,
		MPCost: 50,
		unit:   unit,
	}
}

func (a *SelfHealing) IsMpEnough() bool {
	if a.unit.GetMp() < a.MPCost {
		return false
	}
	return true
}

func (a *SelfHealing) GetMPCost() int {
	return a.MPCost
}

func (a *SelfHealing) SelectTarget([]IUnit) {
}

func (a *SelfHealing) Do() {
	a.unit.TakeDamage(a.Damage)
}
