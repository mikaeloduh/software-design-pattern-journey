package entity

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
