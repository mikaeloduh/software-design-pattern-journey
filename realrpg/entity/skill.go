package entity

import (
	"fmt"
	"io"
	"os"
)

type ISkill interface {
	IObserver
	IsMpEnough() bool
	SelectTarget(targets ...IUnit) error
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

func (a *BasicAttack) SelectTarget(targets ...IUnit) error {
	if len(targets) != 1 {
		return fmt.Errorf("invalid number of args: need 1")
	}
	a.targets = targets
	return nil
}

func (a *BasicAttack) Do() {
	damage := a.unit.GetState().OnAttack(a.unit.GetSTR())
	for _, target := range a.targets {
		target.TakeDamage(damage)
	}
}

func (a *BasicAttack) Update(_ IObservable) {
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

func (a *WaterBall) SelectTarget(targets ...IUnit) error {
	if len(targets) != 1 {
		return fmt.Errorf("invalid number of args: need 1")
	}
	a.targets = targets
	return nil
}

func (a *WaterBall) Do() {
	for _, unit := range a.targets {
		unit.SetHp(unit.GetHp() - a.Damage)
	}

	a.unit.ConsumeMp(a.MPCost)
}

func (a *WaterBall) Update(_ IObservable) {
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

func (a *SelfExplosion) SelectTarget(targets ...IUnit) error {
	a.targets = targets
	return nil
}

func (a *SelfExplosion) Do() {
	for _, target := range a.targets {
		target.TakeDamage(a.Damage)
	}

	a.unit.SetHp(0)
}

func (a *SelfExplosion) Update(_ IObservable) {
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

func (a *CheerUp) SelectTarget(targets ...IUnit) error {
	if len(targets) > 3 {
		return fmt.Errorf("invalid number of args: need 3 or less")
	}
	a.targets = targets
	return nil
}

func (a *CheerUp) Do() {
	for _, target := range a.targets {
		target.SetState(NewCheerUpState(target))
	}
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

func (a *SelfHealing) SelectTarget(_ ...IUnit) error {
	return nil
}

func (a *SelfHealing) Do() {
	a.unit.TakeDamage(a.Damage)
}

func (a *SelfHealing) Update(_ IObservable) {
}

// Summon
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

func (a *Summon) SelectTarget(_ ...IUnit) error {
	return nil
}

func (a *Summon) Do() {
	slime := NewSlime(a.Writer)
	a.unit.GetTroop().AddUnit(slime)

	slime.Register(a)
}

func (a *Summon) Update(_ IObservable) {
	a.unit.TakeDamage(-30)
}

// Curse
type Curse struct {
	Damage  int
	MPCost  int
	unit    IUnit
	targets []IUnit
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

func (a *Curse) SelectTarget(targets ...IUnit) error {
	if len(targets) != 1 {
		return fmt.Errorf("invalid number of args: need 1")
	}
	a.targets = targets
	return nil
}

func (a *Curse) Do() {
	a.targets[0].Register(a)
}

func (a *Curse) Update(_ IObservable) {
	a.unit.TakeDamage(-a.targets[0].GetMp())
}

type OnePunch struct {
	Damage  int
	MPCost  int
	unit    IUnit
	targets []IUnit
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

func (a *OnePunch) SelectTarget(targets ...IUnit) error {
	if len(targets) != 1 {
		return fmt.Errorf("invalid number of args: need 1")
	}
	a.targets = targets
	return nil
}

func (a *OnePunch) Do() {
	a.handler.Do(a.targets[0])
}

func (a *OnePunch) Update(_ IObservable) {
}
