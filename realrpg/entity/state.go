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

func (s *DeadState) OnAttack(_ int) int {
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

func (s *CheerUpState) OnAttack(damage int) int {
	return damage + 50
}

func (s *CheerUpState) OnRoundStart() {
	s.lifetime--
	if s.lifetime <= 0 {
		s.unit.SetState(NewNormalState(s.unit))
	}
}

// PetrochemicalState
type PetrochemicalState struct {
	unit     IUnit
	lifetime int
}

func NewPetrochemicalState(unit IUnit) *PetrochemicalState {
	return &PetrochemicalState{
		unit:     unit,
		lifetime: 3,
	}
}

func (s *PetrochemicalState) OnAttack(damage int) int {
	return damage
}

func (s *PetrochemicalState) OnRoundStart() {
	s.lifetime--
	if s.lifetime <= 0 {
		s.unit.SetState(NewNormalState(s.unit))
	}
}

// PoisonedState
type PoisonedState struct {
	unit     IUnit
	lifetime int
}

func NewPoisonedState(unit IUnit) *PoisonedState {
	return &PoisonedState{
		unit:     unit,
		lifetime: 3,
	}
}

func (s *PoisonedState) OnAttack(damage int) int {
	return damage
}

func (s *PoisonedState) OnRoundStart() {
	s.unit.TakeDamage(30)

	s.lifetime--
	if s.lifetime <= 0 {
		s.unit.SetState(NewNormalState(s.unit))
	}
}
