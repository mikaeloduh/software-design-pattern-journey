package entity

// InvincibleState
type InvincibleState struct {
	character IMapObject
	lifetime  Round
}

func NewInvincibleState(character IMapObject) *InvincibleState {
	return &InvincibleState{character: character, lifetime: 2}
}

func (s *InvincibleState) OnRoundStart() {
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *InvincibleState) OnTakeDamage(_ int) int {
	return 0
}

func (s *InvincibleState) OnAttack(attack AttackMap) AttackMap {
	return attack
}
