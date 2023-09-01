package entity

type EruptingState struct {
	character IMapObject
	lifetime  Round
}

func NewEruptingState(character IMapObject) *EruptingState {
	return &EruptingState{character: character, lifetime: 3}
}

func (s *EruptingState) OnRoundStart() {
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewTeleportState(s.character))
	}
}

func (s *EruptingState) OnTakeDamage(damage int) int {
	return damage
}

func (s *EruptingState) OnAttack(_ AttackMap) AttackMap {
	var a AttackMap
	for y := 0; y <= 9; y++ {
		for x := 0; x <= 9; x++ {
			if !(x == s.character.GetPosition().X && y == s.character.GetPosition().Y) {
				a[y][x] = 50
			}
		}
	}
	return a
}
