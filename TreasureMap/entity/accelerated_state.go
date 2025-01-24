package entity

// AcceleratedState
type AcceleratedState struct {
	character IStatefulMapObject
	lifetime  Round
}

func NewAcceleratedState(character IStatefulMapObject) *AcceleratedState {
	character.SetSpeed(2)
	return &AcceleratedState{character: character, lifetime: 3}
}

func (s *AcceleratedState) OnRoundStart() {
	s.character.SetSpeed(2)
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *AcceleratedState) OnTakeDamage(damage Damage) Damage {
	s.character.SetSpeed(1)
	s.character.SetState(NewNormalState(s.character))
	return damage
}

func (s *AcceleratedState) OnAttack(attack IAttackStrategy) IAttackStrategy {
	return attack
}
