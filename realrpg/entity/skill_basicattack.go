package entity

import "fmt"

// BasicAttack is a built-in skill
type BasicAttack struct {
	unit IUnit
}

func NewBasicAttack(unit IUnit) *BasicAttack {
	return &BasicAttack{unit: unit}
}

func (a *BasicAttack) IsMpEnough() bool {
	return true
}

func (a *BasicAttack) BeforeDo(targets ...IUnit) error {
	if len(targets) != 1 {
		return fmt.Errorf("invalid number of args: need 1")
	}

	return nil
}

func (a *BasicAttack) Do(targets ...IUnit) error {
	if err := a.BeforeDo(targets...); err != nil {
		return err
	}

	damage := a.unit.GetState().OnAttack(a.unit.GetSTR())
	for _, target := range targets {
		target.TakeDamage(damage)
	}

	return nil
}

func (a *BasicAttack) Update(_ IObservable) {
}
