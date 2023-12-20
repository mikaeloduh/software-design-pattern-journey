package entity

// SelfExplosion
type SelfExplosion struct {
	MPCost int
	Damage int
	unit   IUnit
}

func NewSelfExplosion(unit IUnit) *SelfExplosion {
	return &SelfExplosion{
		MPCost: 200,
		Damage: 150,
		unit:   unit,
	}
}

func (a *SelfExplosion) IsMpEnough() bool {
	if a.unit.GetMp() < a.MPCost {
		return false
	}

	return true
}

func (a *SelfExplosion) BeforeDo(targets ...IUnit) error {
	return nil
}

func (a *SelfExplosion) Do(targets ...IUnit) error {
	damage := a.unit.GetState().OnAttack(a.Damage)
	for _, target := range targets {
		target.TakeDamage(damage)
	}

	a.unit.TakeDamage(99999)

	return nil
}

func (a *SelfExplosion) Update(_ IObservable) {
}
