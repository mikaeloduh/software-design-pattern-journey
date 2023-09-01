package entity

import "treasuremap/utils"

type OrderlessState struct {
	character IMapObject
	lifetime  Round
}

func NewOrderlessState(character IMapObject) *OrderlessState {
	return &OrderlessState{character: character, lifetime: 3}
}

func (s OrderlessState) OnRoundStart() {
	if utils.RandBool() {
		//s.character.DisableAction("MoveUp")
		//s.character.DisableAction("MoveDown")
	} else {
		//s.character.DisableAction("MoveRight")
		//s.character.DisableAction("MoveLeft")
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
