package entity

type HumanPlayer struct {
	id        int
	name      string
	HandCards []Card
}

func NewHumanPlayer(id int) *HumanPlayer {
	return &HumanPlayer{id: id}
}

func (p *HumanPlayer) Id() int {
	return p.id
}

func (p *HumanPlayer) Name() string {
	return p.name
}

func (p *HumanPlayer) SetName(name string) {
	p.name = name
}

func (p *HumanPlayer) GetDrawCard(deck *Deck) {
	p.HandCards = append(p.HandCards, deck.DrawCard())
}
