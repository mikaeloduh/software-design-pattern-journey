package entity

import "fmt"

// WaterBall
type WaterBall struct {
	Damage int
	MPCost int
	unit   IUnit
}

func NewWaterBall(unit IUnit) *WaterBall {
	return &WaterBall{
		Damage: 120,
		MPCost: 50,
		unit:   unit,
	}
}

func (a *WaterBall) IsMpEnough() bool {
	if a.unit.GetMp() < a.MPCost {
		return false
	}

	return true
}

func (a *WaterBall) BeforeDo(targets ...IUnit) error {
	if len(targets) != 1 {
		return fmt.Errorf("invalid number of args: need 1")
	}

	return nil
}

func (a *WaterBall) Do(targets ...IUnit) error {
	if err := a.BeforeDo(targets...); err != nil {
		return err
	}

	damage := a.unit.GetState().OnAttack(a.Damage)
	for _, target := range targets {
		target.TakeDamage(damage)
	}

	a.unit.ConsumeMp(a.MPCost)

	return nil
}

func (a *WaterBall) Update(_ IObservable) {
}
