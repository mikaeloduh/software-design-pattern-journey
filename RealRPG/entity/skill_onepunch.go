package entity

import "fmt"

// OnePunch
type OnePunch struct {
	Damage  int
	MPCost  int
	unit    IUnit
	handler ISkillHandler
}

func NewOnePunch(unit IUnit, handler ISkillHandler) *OnePunch {
	return &OnePunch{
		MPCost:  180,
		unit:    unit,
		handler: handler,
	}
}

func (a *OnePunch) IsMpEnough() bool {
	if a.unit.GetMp() < a.MPCost {
		return false
	}

	return true
}

func (a *OnePunch) BeforeDo(targets ...IUnit) error {
	if len(targets) != 1 {
		return fmt.Errorf("invalid number of args: need 1")
	}

	return nil
}

func (a *OnePunch) Do(targets ...IUnit) error {
	if err := a.BeforeDo(targets...); err != nil {
		return err
	}

	a.handler.Do(targets[0], a.unit)

	return nil
}

func (a *OnePunch) Update(IObservable) {}
