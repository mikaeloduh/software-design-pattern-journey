package entity

// NormalState
type NormalState struct {
	Unit IUnit
}

func NewNormalState(unit IUnit) *NormalState {
	return &NormalState{Unit: unit}
}

func (s *NormalState) OnAttack(damage int) int {
	return damage
}

func (s *NormalState) OnRoundStart() {
}
