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

func (s *NormalState) OnTakeDamage(d int) int {
	return d
}

func (s *NormalState) OnAttack(attack IAttackStrategy) IAttackStrategy {
	return attack
}
