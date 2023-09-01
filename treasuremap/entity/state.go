package entity

import (
	"treasuremap/commons"
)

// IState interface
type IState interface {
	OnRoundStart()
	OnTakeDamage(damage int) int
	OnAttack(attack AttackMap) AttackMap
}

// NormalState
type NormalState struct {
	character *Character
}

func NewNormalState(character *Character) *NormalState {
	return &NormalState{character: character}
}

func (s *NormalState) OnRoundStart() {
}

func (s *NormalState) OnTakeDamage(d int) int {
	return d
}

func (s *NormalState) OnAttack(attack AttackMap) AttackMap {
	return attack
}

// PoisonedState
type PoisonedState struct {
	character *Character
	lifetime  Round
}

func NewPoisonedState(character *Character) *PoisonedState {
	return &PoisonedState{character: character, lifetime: 3}
}

func (s *PoisonedState) OnRoundStart() {
	s.character.TakeDamage(15)
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *PoisonedState) OnTakeDamage(d int) int {
	return d
}

func (s *PoisonedState) OnAttack(attack AttackMap) AttackMap {
	return attack
}

// InvincibleState
type InvincibleState struct {
	character *Character
	lifetime  Round
}

func NewInvincibleState(character *Character) *InvincibleState {
	return &InvincibleState{character: character, lifetime: 2}
}

func (s *InvincibleState) OnRoundStart() {
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *InvincibleState) OnTakeDamage(_ int) int {
	return 0
}

func (s *InvincibleState) OnAttack(attack AttackMap) AttackMap {
	return attack
}

// HealingState
type HealingState struct {
	character *Character
	lifetime  Round
}

func NewHealingState(character *Character) *HealingState {
	return &HealingState{character: character, lifetime: 5}
}

func (s *HealingState) OnRoundStart() {
	s.character.Heal(30)
	s.lifetime--
	if s.lifetime <= 0 || s.character.Hp >= s.character.MaxHp {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *HealingState) OnTakeDamage(d int) int {
	return d
}

func (s *HealingState) OnAttack(attack AttackMap) AttackMap {
	return attack
}

// AcceleratedState
type AcceleratedState struct {
	character *Character
	lifetime  Round
}

func NewAcceleratedState(character *Character) *AcceleratedState {
	character.SetSpeed(2)
	return &AcceleratedState{character: character, lifetime: 3}
}

func (s *AcceleratedState) OnRoundStart() {
	s.character.SetSpeed(2)
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *AcceleratedState) OnTakeDamage(damage int) int {
	s.character.SetSpeed(1)
	s.character.SetState(NewNormalState(s.character))
	return damage
}

func (s *AcceleratedState) OnAttack(attack AttackMap) AttackMap {
	return attack
}

type OrderlessState struct {
	character *Character
	lifetime  Round
}

func NewOrderlessState(character *Character) *OrderlessState {
	return &OrderlessState{character: character, lifetime: 3}
}

func (s OrderlessState) OnRoundStart() {
	if commons.RandBool() {
		s.character.DisableAction("MoveUp")
		s.character.DisableAction("MoveDown")
	} else {
		s.character.DisableAction("MoveRight")
		s.character.DisableAction("MoveLeft")
	}

	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s OrderlessState) OnTakeDamage(damage int) int {
	return damage
}

func (s OrderlessState) OnAttack(attack AttackMap) AttackMap {
	panic("operation not allowed")
}

// StockpileState
type StockpileState struct {
	character *Character
	lifetime  Round
}

func NewStockpileState(character *Character) *StockpileState {
	return &StockpileState{character: character, lifetime: 2}
}

func (s *StockpileState) OnRoundStart() {
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewEruptingState(s.character))
	}
}

func (s *StockpileState) OnTakeDamage(damage int) int {
	return damage
}

func (s *StockpileState) OnAttack(attack AttackMap) AttackMap {
	return attack
}

type EruptingState struct {
	character *Character
	lifetime  Round
}

func NewEruptingState(character *Character) *EruptingState {
	return &EruptingState{character: character, lifetime: 3}
}

func (s *EruptingState) OnRoundStart() {
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *EruptingState) OnTakeDamage(damage int) int {
	return damage
}

func (s *EruptingState) OnAttack(_ AttackMap) AttackMap {
	var a AttackMap
	for y := 0; y <= 9; y++ {
		for x := 0; x <= 9; x++ {
			if !(x == s.character.Position.X && y == s.character.Position.Y) {
				a[y][x] = 50
			}
		}
	}
	return a
}
