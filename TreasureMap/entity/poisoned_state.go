package entity

// PoisonedState
type PoisonedState struct {
	character IStatefulMapObject
	lifetime  Round
}

func NewPoisonedState(character IStatefulMapObject) *PoisonedState {
	return &PoisonedState{character: character, lifetime: 3}
}

func (s *PoisonedState) OnRoundStart() {
	s.character.TakeDamage(15)
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *PoisonedState) OnTakeDamage(damage Damage) Damage {
	return damage
}

func (s *PoisonedState) OnAttack(attack IAttackStrategy) IAttackStrategy {
	return attack
}
