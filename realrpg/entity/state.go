package entity

type IState interface {
}

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

// CheerUpState
type CheerUpState struct {
	unit IUnit
}

func NewCheerUpState(unit IUnit) *CheerUpState {
	return &CheerUpState{unit: unit}
}

func (c *CheerUpState) OnAttack(damage int) int {
	return damage + 50
}
