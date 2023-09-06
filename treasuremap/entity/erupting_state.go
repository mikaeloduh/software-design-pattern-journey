package entity

type EruptingState struct {
	character IStatefulMapObject
	lifetime  Round
}

func NewEruptingState(character IStatefulMapObject) *EruptingState {
	return &EruptingState{character: character, lifetime: 3}
}

func (s *EruptingState) OnRoundStart() {
	s.lifetime--
	if s.lifetime <= 0 {
		s.character.SetState(NewTeleportState(s.character))
	}
}

func (s *EruptingState) OnTakeDamage(damage Damage) Damage {
	return damage
}

func (s *EruptingState) OnAttack(_ IAttackStrategy) IAttackStrategy {

	return func(worldMap [10][10]*Position) (damageArea AttackDamageArea) {
		// Attack hits every character on the entire map, dealing 50 damage with each attack
		for y := 0; y < len(worldMap); y++ {
			for x := 0; x < len(worldMap[0]); x++ {
				if !(x == s.character.GetPosition().X && y == s.character.GetPosition().Y) {
					damageArea[y][x] = 50
				}
			}
		}
		return
	}
}
