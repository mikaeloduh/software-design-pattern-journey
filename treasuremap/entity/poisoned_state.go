package entity

// PoisonedState
type PoisonedState struct {
	character IMapObject
	lifetime  Round
}

func NewPoisonedState(character IMapObject) *PoisonedState {
	return &PoisonedState{character: character, lifetime: 3}
}

func (s *PoisonedState) OnRoundStart() {
	s.character.TakeDamage(15)
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *PoisonedState) OnTakeDamage(d int) int {
	return d
}

func (s *PoisonedState) OnAttack(attack AttackMap) AttackMap {
	return attack
}
