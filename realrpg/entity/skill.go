package entity

type ISkill interface {
	SelectTarget([]IUnit)
	Do()
	GetMPCost() int
}

// BasicAttack
type BasicAttack struct {
	unit    IUnit
	targets []IUnit
}

func NewBasicAttack(unit IUnit) *BasicAttack {
	return &BasicAttack{unit: unit}
}

func (a *BasicAttack) SelectTarget(units []IUnit) {
	a.targets = units
}

func (a *BasicAttack) Do() {
	damage := a.unit.GetState().OnAttack(a.unit.GetSTR())
	for _, unit := range a.targets {
		unit.SetHp(unit.GetHp() - damage)
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

type SelfExplosion struct {
	MPCost  int
	Damage  int
	unit    IUnit
	targets []IUnit
}

func NewSelfExplosion(unit IUnit) *SelfExplosion {
	return &SelfExplosion{
		MPCost: 200,
		Damage: 150,
		unit:   unit,
	}
}

func (s *SelfExplosion) SelectTarget(unit []IUnit) {
	s.targets = unit
}

func (s *SelfExplosion) Do() {
	for _, target := range s.targets {
		target.SetHp(target.GetHp() - s.Damage)
	}

	s.unit.SetHp(0)
}

func (s *SelfExplosion) GetMPCost() int {
	return s.MPCost
}

type CheerUp struct {
	Damage  int
	MPCost  int
	unit    IUnit
	targets []IUnit
}

func (s *CheerUp) GetMPCost() int {
	return s.MPCost
}

func (s *CheerUp) SelectTarget(units []IUnit) {
	s.targets = units
}

func (s *CheerUp) Do() {
	for _, unit := range s.targets {
		unit.SetState(NewCheerUpState(unit))
	}
}

type SelfHealing struct {
	Damage  int
	MPCost  int
	targets []IUnit
}

func (s *SelfHealing) GetMPCost() int {
	return s.MPCost
}

func (s *SelfHealing) SelectTarget(units []IUnit) {
	s.targets = units
}

func (s *SelfHealing) Do() {
	for _, unit := range s.targets {
		unit.SetHp(unit.GetHp() + s.Damage)
	}
}
