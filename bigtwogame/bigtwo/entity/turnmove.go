package entity

type TurnMove struct {
	hand     *[]BigTwoCard
	selected int
}

func NewTurnMove(hand *[]BigTwoCard, selected int) *TurnMove {
	return &TurnMove{hand: hand, selected: selected}
}

func (t *TurnMove) DryRun() BigTwoCard {
	return (*t.hand)[t.selected]
}

func (t *TurnMove) Play() BigTwoCard {
	card := (*t.hand)[t.selected]
	*t.hand = append((*t.hand)[:t.selected], (*t.hand)[t.selected+1:]...)
	return card
}
