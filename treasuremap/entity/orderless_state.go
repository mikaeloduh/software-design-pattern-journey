package entity

import "treasuremap/utils"

type OrderlessState struct {
	character IStatefulMapObject
	lifetime  Round
}

func NewOrderlessState(character IStatefulMapObject) *OrderlessState {
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

func (s OrderlessState) OnTakeDamage(damage Damage) Damage {
	return damage
}

func (s OrderlessState) OnAttack(attack IAttackStrategy) IAttackStrategy {
	panic("operation not allowed")
}
