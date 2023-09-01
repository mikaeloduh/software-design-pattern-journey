package entity

import "math/rand"

// TeleportState
type TeleportState struct {
	character IMapObject
	lifetime  Round
}

func NewTeleportState(character IMapObject) *TeleportState {
	return &TeleportState{character: character, lifetime: 1}
}

func (s *TeleportState) OnRoundStart() {
	s.lifetime--
	if s.lifetime <= 0 {
		x, y := rand.Intn(10), rand.Intn(10)
		s.character.GetPosition().Move(x, y, s.character.GetPosition().Direction)

		s.character.SetState(NewNormalState(s.character))
	}
}

func (s *TeleportState) OnTakeDamage(damage int) int {
	return damage
}

func (s *TeleportState) OnAttack(attack AttackMap) AttackMap {
	return attack
}
