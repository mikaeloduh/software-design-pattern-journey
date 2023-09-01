package entity

// StockpileState
type StockpileState struct {
	character IMapObject
	lifetime  Round
}

func NewStockpileState(character IMapObject) *StockpileState {
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
