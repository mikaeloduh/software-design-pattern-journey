package entity

// StockpileState
type StockpileState struct {
	character IStatefulMapObject
	lifetime  Round
}

func NewStockpileState(character IStatefulMapObject) *StockpileState {
	return &StockpileState{character: character, lifetime: 2}
}

func (s *StockpileState) OnRoundStart() {
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewEruptingState(s.character))
	}
}

func (s *StockpileState) OnTakeDamage(damage Damage) Damage {
	return damage
}

func (s *StockpileState) OnAttack(attack IAttackStrategy) IAttackStrategy {
	return attack
}
