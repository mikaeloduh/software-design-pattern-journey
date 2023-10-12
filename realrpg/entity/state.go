package entity

type IState interface {
	OnAttack(damage int) int
	OnRoundStart()
}

type DeadState struct {
	Unit IUnit
}

func NewDeadState(unit IUnit) *DeadState {
	unit.SetHp(0)
	unit.Notify()

	return &DeadState{Unit: unit}
}

func (s *DeadState) OnAttack(damage int) int {
	return 0
}

func (s *DeadState) OnRoundStart() {
}

// NormalState
type NormalState struct {
	Unit IUnit
}

func NewNormalState(unit IUnit) *NormalState {
	return &NormalState{Unit: unit}
}

func (s *NormalState) OnAttack(damage int) int {
	return damage
}

func (s *NormalState) OnRoundStart() {
}

// CheerUpState
type CheerUpState struct {
	unit     IUnit
	lifetime int
}

func NewCheerUpState(unit IUnit) *CheerUpState {
	return &CheerUpState{unit: unit, lifetime: 3}
}

func (s *CheerUpState) OnRoundStart() {
	s.lifetime--
	if s.lifetime <= 0 {
		s.unit.SetState(NewNormalState(s.unit))
	}
}

func (s *CheerUpState) OnAttack(damage int) int {
	return damage + 50
}
