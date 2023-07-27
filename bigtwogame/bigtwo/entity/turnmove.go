package entity

type TurnMove struct {
	hand     *[]BigTwoCard
	selected []BigTwoCard
}

func NewTurnMove(hand *[]BigTwoCard, selected []BigTwoCard) *TurnMove {
	return &TurnMove{hand: hand, selected: selected}
}

func (t *TurnMove) DryRun() []BigTwoCard {
	return t.selected
}

func (t *TurnMove) Play() []BigTwoCard {
	cardsToRemove := make(map[BigTwoCard]bool)
	for _, c := range t.selected {
		cardsToRemove[c] = true
	}

	var filteredCards []BigTwoCard
	for _, c := range *t.hand {
		if _, found := cardsToRemove[c]; !found {
			filteredCards = append(filteredCards, c)
		}

	}
	*t.hand = filteredCards

	return t.selected
}
