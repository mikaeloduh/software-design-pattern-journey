package entity

// HealingState
type HealingState struct {
	character IStatefulMapObject
	lifetime  Round
}

func NewHealingState(character IStatefulMapObject) *HealingState {
	return &HealingState{character: character, lifetime: 5}
}

func (s *HealingState) OnRoundStart() {
	s.character.Heal(30)
	s.lifetime--
	if s.lifetime <= 0 || s.character.GetHp() >= s.character.GetMaxHp() {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *HealingState) OnTakeDamage(damage Damage) Damage {
	return damage
}

func (s *HealingState) OnAttack(attack IAttackStrategy) IAttackStrategy {
	return attack
}
