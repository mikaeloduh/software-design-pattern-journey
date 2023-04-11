package entity

type AIPlayer struct {
	id        int
	name      string
	HandCards []Card
}

func (ai *AIPlayer) TakeTurn(deck *Deck) {
	// 1. exchange?

	// 2. show
	play := 0
	deck.Table[ai.id] = ai.HandCards[play]
	ai.HandCards = append([]Card{}, append(ai.HandCards[0:play], ai.HandCards[play+1:]...)...)
}

func (ai *AIPlayer) Id() int {
	return ai.id
}

func (ai *AIPlayer) Name() string {
	return ai.name
}

func (ai *AIPlayer) GetDrawCard(deck *Deck) {
	ai.HandCards = append(ai.HandCards, deck.DrawCard())
}

func NewAIPlayer(id int) *AIPlayer {
	return &AIPlayer{
		id:   id,
		name: "AI has no name",
	}
}
