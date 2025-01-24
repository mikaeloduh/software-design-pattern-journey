package entity

type DeadState struct {
	Unit IUnit
}

func NewDeadState(unit IUnit) *DeadState {
	unit.SetHp(0)
	unit.Notify()

	return &DeadState{Unit: unit}
}

func (s *DeadState) OnAttack(_ int) int {
	return 0
}

func (s *DeadState) OnRoundStart() {
}
