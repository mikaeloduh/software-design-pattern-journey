package entity

// NormalState
type NormalState struct {
	character IStatefulMapObject
}

func NewNormalState(character IStatefulMapObject) *NormalState {
	return &NormalState{character: character}
}

func (s *NormalState) OnRoundStart() {
}

func (s *NormalState) OnTakeDamage(damage Damage) Damage {
	return damage
}

func (s *NormalState) OnAttack(attack IAttackStrategy) IAttackStrategy {
	return attack
}
