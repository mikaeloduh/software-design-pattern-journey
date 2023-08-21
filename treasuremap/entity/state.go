package entity

// IState interface
type IState interface {
	OnRoundStart()
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

// PoisonedState
type PoisonedState struct {
	character *Character
}

func NewPoisonedState(character *Character) *PoisonedState {
	return &PoisonedState{character: character}
}

func (s *PoisonedState) OnRoundStart() {
	s.character.AddHp(-15)
}
