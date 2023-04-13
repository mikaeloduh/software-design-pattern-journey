package entity

type HumanPlayer struct {
	id        int
	name      string
	HandCards []Card
	point     int
}

func (p *HumanPlayer) YouExchangeMyCard(card Card) Card {
	// TODO: Choose a card
	myCard := p.HandCards[0]
	p.HandCards[0] = card

	return myCard
}

func (p *HumanPlayer) MeExchangeYourCard(player IPlayer) {
	// TODO: Choose a card
	c := p.HandCards[0]

	// Exchange it
	ex := player.YouExchangeMyCard(c)
	p.HandCards[0] = ex
}

func (p *HumanPlayer) Point() int {
	return p.point
}

func (p *HumanPlayer) AddPoint() {
	p.point += 1
}

func (p *HumanPlayer) TakeTurn() Card {
	// 1. exchange?

	// 2. show
	play := 0
	showCard := p.HandCards[play]
	p.HandCards = append([]Card{}, append(p.HandCards[0:play], p.HandCards[play+1:]...)...)

	return showCard
}

func (p *HumanPlayer) GetCard(card Card) {
	p.HandCards = append(p.HandCards, card)
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

func NewHumanPlayer(id int) *HumanPlayer {
	return &HumanPlayer{id: id}
}
