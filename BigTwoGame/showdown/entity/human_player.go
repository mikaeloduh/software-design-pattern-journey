package entity

type HumanPlayer struct {
	id    int
	Name  string
	Hand  []ShowDownCard
	point int
	count int
	IPlayerInput
	IPlayerOutput
}

func (p *HumanPlayer) GetName() string {
	return p.Name
}

func (p *HumanPlayer) GetHand() []ShowDownCard {
	return p.Hand
}

func NewHumanPlayer(input IPlayerInput, output IPlayerOutput) *HumanPlayer {
	return &HumanPlayer{
		count:         3,
		IPlayerInput:  input,
		IPlayerOutput: output,
	}
}

func (p *HumanPlayer) TakeTurn() ShowDownCard {
	p.TakeTurnStartOutput(p.Name)
	p.PrintCardsOutput(p.Hand)

	// 2. show
	p.AskShowCardOutput(p.Name)
	toPlay := p.InputNum(0, len(p.Hand)-1)
	showCard := p.Hand[toPlay]
	p.RemoveCard(toPlay)

	return showCard
}

func (p *HumanPlayer) SetCard(card ShowDownCard) {
	p.Hand = append(p.Hand, card)
}

func (p *HumanPlayer) RemoveCard(index int) ShowDownCard {
	card := p.Hand[index]
	p.Hand = append(p.Hand[:index], p.Hand[index+1:]...)
	return card
}

func (p *HumanPlayer) Rename() {
	p.RenameOutput(p.Name)

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

func (p *HumanPlayer) SetName(name string) {
	p.Name = name
}
