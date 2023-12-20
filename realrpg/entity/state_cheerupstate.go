package entity

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
