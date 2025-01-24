package entity

// InvincibleState
type InvincibleState struct {
	character IStatefulMapObject
	lifetime  Round
}

func NewInvincibleState(character IStatefulMapObject) *InvincibleState {
	return &InvincibleState{character: character, lifetime: 2}
}

func (s *InvincibleState) OnRoundStart() {
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *InvincibleState) OnTakeDamage(damage Damage) Damage {
	return 0
}

func (s *InvincibleState) OnAttack(attack IAttackStrategy) IAttackStrategy {
	return attack
}
