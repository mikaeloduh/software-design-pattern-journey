package entity

type TurnMove struct {
	hand     *[]BigTwoCard
	selected []int
}

func NewTurnMove(hand *[]BigTwoCard, selected []int) *TurnMove {
	return &TurnMove{hand: hand, selected: selected}
}

func (t *TurnMove) DryRun() []BigTwoCard {
	var cards []BigTwoCard
	for _, v := range t.selected {
		cards = append(cards, (*t.hand)[v])
	}
	return cards
}

func (t *TurnMove) Play() []BigTwoCard {
	playCards := t.DryRun()
	for i := range t.selected {
		*t.hand = append((*t.hand)[:i], (*t.hand)[i+1:]...)
	}

	return playCards
}
