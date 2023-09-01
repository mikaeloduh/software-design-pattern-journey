package entity

// HealingState
type HealingState struct {
	character IMapObject
	lifetime  Round
}

func NewHealingState(character IMapObject) *HealingState {
	return &HealingState{character: character, lifetime: 5}
}

func (s *HealingState) OnRoundStart() {
	s.character.Heal(30)
	s.lifetime--
	if s.lifetime <= 0 || s.character.GetHp() >= s.character.GetMaxHp() {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *HealingState) OnTakeDamage(d int) int {
	return d
}

func (s *HealingState) OnAttack(attack AttackMap) AttackMap {
	return attack
}
