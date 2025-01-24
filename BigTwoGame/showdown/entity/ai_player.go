package entity

type AIPlayer struct {
	id    int
	Name  string
	Hand  []ShowDownCard
	point int
	count int
	IPlayerInput
	IPlayerOutput
}

func (a *AIPlayer) GetName() string {
	return a.Name
}

func (a *AIPlayer) GetHand() []ShowDownCard {
	return a.Hand
}

func NewAIPlayer(input IPlayerInput, output IPlayerOutput) *AIPlayer {
	return &AIPlayer{
		count:         3,
		Name:          "PlayerAI",
		IPlayerInput:  input,
		IPlayerOutput: output,
	}
}

func (a *AIPlayer) TakeTurn() ShowDownCard {
	a.TakeTurnStartOutput(a.Name)

	// 2. Show card
	a.AskShowCardOutput(a.Name)
	toPlay := a.InputNum(0, len(a.Hand)-1)
	showCard := a.Hand[toPlay]
	a.RemoveCard(toPlay)

	return showCard
}

func (a *AIPlayer) SetCard(card ShowDownCard) {
	a.Hand = append(a.Hand, card)
}

func (a *AIPlayer) RemoveCard(index int) ShowDownCard {
	card := a.Hand[index]
	a.Hand = append(a.Hand[:index], a.Hand[index+1:]...)
	return card
}

func (a *AIPlayer) Rename() {
}

func (a *AIPlayer) Id() int {
	return a.id
}

func (a *AIPlayer) SetId(i int) {
	a.id = i
}

func (a *AIPlayer) Point() int {
	return a.point
}

func (a *AIPlayer) AddPoint() {
	a.point += 1
}

func (a *AIPlayer) SetName(s string) {
	a.Name = s + "_AI"
}
