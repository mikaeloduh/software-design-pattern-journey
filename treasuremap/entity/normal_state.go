package entity

// NormalState
type NormalState struct {
	character IMapObject
}

func NewNormalState(character IMapObject) *NormalState {
	return &NormalState{character: character}
}

func (s *NormalState) OnRoundStart() {
}

func (s *NormalState) OnTakeDamage(d int) int {
	return d
}

func (s *NormalState) OnAttack(attack AttackMap) AttackMap {
	return attack
}
