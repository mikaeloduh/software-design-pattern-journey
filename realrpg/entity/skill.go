package entity

import (
	"io"
	"os"
)

type ISkill interface {
	IsMpEnough() bool
	SelectTarget([]IUnit)
	Do()
}

// BasicAttack
type BasicAttack struct {
	unit    IUnit
	targets []IUnit
}

func NewBasicAttack(unit IUnit) *BasicAttack {
	return &BasicAttack{unit: unit}
}

func (a *BasicAttack) IsMpEnough() bool {
	return true
}

func (a *BasicAttack) SelectTarget(units []IUnit) {
	a.targets = units
}

func (a *BasicAttack) Do() {
	damage := a.unit.GetState().OnAttack(a.unit.GetSTR())
	for _, target := range a.targets {
		target.TakeDamage(damage)
	}
}

// WaterBall
type WaterBall struct {
	Damage  int
	MPCost  int
	unit    IUnit
	targets []IUnit
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

func (a *WaterBall) SelectTarget(units []IUnit) {
	a.targets = units
}

func (a *WaterBall) Do() {
	for _, unit := range a.targets {
		unit.SetHp(unit.GetHp() - a.Damage)
	}

	a.unit.ConsumeMp(a.MPCost)
}

// SelfExplosion
type SelfExplosion struct {
	MPCost  int
	Damage  int
	unit    IUnit
	targets []IUnit
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

func (a *SelfExplosion) SelectTarget(unit []IUnit) {
	a.targets = unit
}

func (a *SelfExplosion) Do() {
	for _, target := range a.targets {
		target.TakeDamage(a.Damage)
	}

	a.unit.SetHp(0)
}

// CheerUp
type CheerUp struct {
	Damage  int
	MPCost  int
	unit    IUnit
	targets []IUnit
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

func (a *CheerUp) SelectTarget(units []IUnit) {
	a.targets = units
}

func (a *CheerUp) Do() {
	for _, target := range a.targets {
		target.SetState(NewCheerUpState(target))
	}
}

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

func (a *SelfHealing) SelectTarget([]IUnit) {
}

func (a *SelfHealing) Do() {
	a.unit.TakeDamage(a.Damage)
}

type Summon struct {
	MPCost int
	unit   IUnit
	Writer io.Writer
}

func NewSummon(unit IUnit) *Summon {
	return &Summon{
		MPCost: 150,
		unit:   unit,
		Writer: os.Stdout,
	}
}

func (a *Summon) IsMpEnough() bool {
	if a.unit.GetMp() < a.MPCost {
		return false
	}
	return true
}

func (a *Summon) SelectTarget(_ []IUnit) {
}

func (a *Summon) Do() {
	slime := NewSlime(a.Writer)

	a.unit.GetTroop().AddUnit(slime)
}
