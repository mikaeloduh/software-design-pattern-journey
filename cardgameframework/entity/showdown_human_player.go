package entity

type ShowdownHumanPlayer struct {
	id        int
	name      string
	HandCards []ShowdownCard
	point     int
	count     int
	IInput
	IOutput
}

func NewHumanPlayer(input IInput, output IOutput) *ShowdownHumanPlayer {
	return &ShowdownHumanPlayer{
		count:   3,
		IInput:  input,
		IOutput: output,
	}
}

func (p *ShowdownHumanPlayer) TakeTurn(players []IPlayer) ShowdownCard {
	p.TakeTurnStartOutput(p.name)
	p.PrintCardsOutput(p.HandCards)

	// 2. show
	p.AskShowCardOutput(p.name)
	toPlay := p.InputNum(0, len(p.HandCards)-1)
	showCard := p.HandCards[toPlay]
	p.HandCards = append([]ShowdownCard{}, append(p.HandCards[0:toPlay], p.HandCards[toPlay+1:]...)...)

	return showCard
}

func (p *ShowdownHumanPlayer) AssignCard(card ShowdownCard) {
	p.HandCards = append(p.HandCards, card)
}

func (p *ShowdownHumanPlayer) Rename() {
	p.RenameOutput(p.name)

	s := p.InputString()
	if s != "?" {
		p.SetName(s)
	}
}

func (p *ShowdownHumanPlayer) Id() int {
	return p.id
}

func (p *ShowdownHumanPlayer) SetId(i int) {
	p.id = i
}

func (p *ShowdownHumanPlayer) Point() int {
	return p.point
}

func (p *ShowdownHumanPlayer) AddPoint() {
	p.point += 1
}

func (p *ShowdownHumanPlayer) Name() string {
	return p.name
}

func (p *ShowdownHumanPlayer) SetName(name string) {
	p.name = name
}
