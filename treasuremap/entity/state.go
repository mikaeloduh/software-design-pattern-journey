package entity

// IState interface
type IState interface {
	OnRoundStart()
	OnTakeDamage(damage int) int
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
