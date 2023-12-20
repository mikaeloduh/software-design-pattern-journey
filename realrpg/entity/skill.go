package entity

import (
	"fmt"
	"io"
	"os"
)

type ISkill interface {
	IObserver
	IsMpEnough() bool
	BeforeDo(targets ...IUnit) error
	Do(targets ...IUnit) error
}

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
