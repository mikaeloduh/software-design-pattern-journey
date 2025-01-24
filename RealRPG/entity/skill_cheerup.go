package entity

import "fmt"

// CheerUp
type CheerUp struct {
	Damage int
	MPCost int
	unit   IUnit
}

func NewCheerUp(unit IUnit) *CheerUp {
	return &CheerUp{
		Damage: 0,
		MPCost: 100,
		unit:   unit,
	}
}

func (a *CheerUp) IsMpEnough() bool {
	if a.unit.GetMp() < a.MPCost {
		return false
	}

	return true
}

func (a *CheerUp) BeforeDo(targets ...IUnit) error {
	if len(targets) > 3 {
		return fmt.Errorf("invalid number of args: need 3 or less")
	}

	return nil
}

func (a *CheerUp) Do(targets ...IUnit) error {
	if err := a.BeforeDo(targets...); err != nil {
		return err
	}

	for _, target := range targets {
		target.SetState(NewCheerUpState(target))
	}

	return nil
}

func (a *CheerUp) Update(_ IObservable) {
}
