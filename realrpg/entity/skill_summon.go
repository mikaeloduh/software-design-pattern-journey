package entity

import (
	"io"
	"os"
)

// Summon
type Summon struct {
	MPCost int
	Writer io.Writer
	unit   IUnit
}

func NewSummon(unit IUnit) *Summon {
	return &Summon{
		MPCost: 150,
		Writer: os.Stdout,
		unit:   unit,
	}
}

func (a *Summon) IsMpEnough() bool {
	if a.unit.GetMp() < a.MPCost {
		return false
	}

	return true
}

func (a *Summon) BeforeDo(...IUnit) error {
	return nil
}

func (a *Summon) Do(...IUnit) error {
	slime := NewSlime(nil, a.Writer)
	slime.Register(a)
	a.unit.GetTroop().AddUnit(slime)

	return nil
}

func (a *Summon) Update(_ IObservable) {
	a.unit.TakeDamage(-30)
}
