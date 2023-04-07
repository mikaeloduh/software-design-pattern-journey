package entity

type AIPlayer struct {
	id        int
	name      string
	HandCards []Card
}

func NewAIPlayer(id int) *AIPlayer {
	return &AIPlayer{
		id:   id,
		name: "AI has no name",
	}
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
