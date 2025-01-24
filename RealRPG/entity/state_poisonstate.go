package entity

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
