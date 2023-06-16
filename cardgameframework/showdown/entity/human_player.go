package entity

type HumanPlayer struct {
	id        int
	name      string
	HandCards []Card
	point     int
	count     int
	IInput
	IOutput
}

func NewHumanPlayer(input IInput, output IOutput) *HumanPlayer {
	return &HumanPlayer{
		count:   3,
		IInput:  input,
		IOutput: output,
	}
}

func (p *HumanPlayer) TakeTurn(players []IPlayer) Card {
	p.TakeTurnStartOutput(p.name)
	p.PrintCardsOutput(p.HandCards)

	// 2. show
	p.AskShowCardOutput(p.name)
	toPlay := p.InputNum(0, len(p.HandCards)-1)
	showCard := p.HandCards[toPlay]
	p.HandCards = append([]Card{}, append(p.HandCards[0:toPlay], p.HandCards[toPlay+1:]...)...)

	return showCard
}

func (p *HumanPlayer) AssignCard(card Card) {
	p.HandCards = append(p.HandCards, card)
}

func (p *HumanPlayer) Rename() {
	p.RenameOutput(p.name)

	s := p.InputString()
	if s != "?" {
		p.SetName(s)
	}
}

func (p *HumanPlayer) Id() int {
	return p.id
}

func (p *HumanPlayer) SetId(i int) {
	p.id = i
}

func (p *HumanPlayer) Point() int {
	return p.point
}

func (p *HumanPlayer) AddPoint() {
	p.point += 1
}

func (p *HumanPlayer) Name() string {
	return p.name
}

func (p *HumanPlayer) SetName(name string) {
	p.name = name
}
