package entity

// SelfHealing
type SelfHealing struct {
	Damage int
	MPCost int
	unit   IUnit
}

func NewSelfHealing(unit IUnit) *SelfHealing {
	return &SelfHealing{
		Damage: -50,
		MPCost: 50,
		unit:   unit,
	}
}

func (a *SelfHealing) IsMpEnough() bool {
	if a.unit.GetMp() < a.MPCost {
		return false
	}

	return true
}

func (a *SelfHealing) BeforeDo(...IUnit) error {
	return nil
}

func (a *SelfHealing) Do(...IUnit) error {
	a.unit.TakeDamage(a.Damage)

	return nil
}

func (a *SelfHealing) Update(_ IObservable) {}
