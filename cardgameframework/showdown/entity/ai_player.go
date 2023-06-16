package entity

type AIPlayer struct {
	id        int
	name      string
	HandCards []Card
	point     int
	count     int
	IInput
	IOutput
}

func NewAIPlayer(input IInput, output IOutput) *AIPlayer {
	return &AIPlayer{
		count:   3,
		name:    "PlayerAI",
		IInput:  input,
		IOutput: output,
	}
}

func (ai *AIPlayer) TakeTurn(players []IPlayer) Card {
	ai.TakeTurnStartOutput(ai.name)

	// 2. Show card
	ai.AskShowCardOutput(ai.name)
	toPlay := ai.InputNum(0, len(ai.HandCards)-1)
	showCard := ai.HandCards[toPlay]
	ai.HandCards = append([]Card{}, append(ai.HandCards[0:toPlay], ai.HandCards[toPlay+1:]...)...)

	return showCard
}

func (ai *AIPlayer) AssignCard(card Card) {
	ai.HandCards = append(ai.HandCards, card)
}

func (ai *AIPlayer) Rename() {
}

func (ai *AIPlayer) Id() int {
	return ai.id
}

func (ai *AIPlayer) SetId(i int) {
	ai.id = i
}

func (ai *AIPlayer) Point() int {
	return ai.point
}

func (ai *AIPlayer) AddPoint() {
	ai.point += 1
}

func (ai *AIPlayer) Name() string {
	return ai.name
}

func (ai *AIPlayer) SetName(s string) {
	ai.name = s + "_AI"
}
