package entity

import "fmt"

// Curse
type Curse struct {
	Damage int
	MPCost int
	unit   IUnit
}

func NewCurse(unit IUnit) *Curse {
	return &Curse{
		Damage: 0,
		MPCost: 100,
		unit:   unit,
	}
}

func (a *Curse) IsMpEnough() bool {
	if a.unit.GetMp() < a.MPCost {
		return false
	}

	return true
}

func (a *Curse) BeforeDo(targets ...IUnit) error {
	if len(targets) != 1 {
		return fmt.Errorf("invalid number of args: need 1")
	}

	return nil
}

func (a *Curse) Do(targets ...IUnit) error {
	if err := a.BeforeDo(targets...); err != nil {
		return err
	}

	targets[0].Register(a)

	return nil
}

func (a *Curse) Update(target IObservable) {
	a.unit.TakeDamage(-target.(IUnit).GetMp())
}
