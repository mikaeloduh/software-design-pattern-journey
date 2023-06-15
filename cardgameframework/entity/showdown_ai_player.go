package entity

type ShowdownAIPlayer struct {
	id        int
	name      string
	HandCards []ShowdownCard
	point     int
	count     int
	IShowdownInput
	IShowdownOutput
}

func NewAIPlayer(input IShowdownInput, output IShowdownOutput) *ShowdownAIPlayer {
	return &ShowdownAIPlayer{
		count:           3,
		name:            "PlayerAI",
		IShowdownInput:  input,
		IShowdownOutput: output,
	}
}

func (ai *ShowdownAIPlayer) TakeTurn(players []IShowdownPlayer) ShowdownCard {
	ai.TakeTurnStartOutput(ai.name)

	// 2. Show card
	ai.AskShowCardOutput(ai.name)
	toPlay := ai.InputNum(0, len(ai.HandCards)-1)
	showCard := ai.HandCards[toPlay]
	ai.HandCards = append([]ShowdownCard{}, append(ai.HandCards[0:toPlay], ai.HandCards[toPlay+1:]...)...)

	return showCard
}

func (ai *ShowdownAIPlayer) AssignCard(card ShowdownCard) {
	ai.HandCards = append(ai.HandCards, card)
}

func (ai *ShowdownAIPlayer) Rename() {
}

func (ai *ShowdownAIPlayer) Id() int {
	return ai.id
}

func (ai *ShowdownAIPlayer) SetId(i int) {
	ai.id = i
}

func (ai *ShowdownAIPlayer) Point() int {
	return ai.point
}

func (ai *ShowdownAIPlayer) AddPoint() {
	ai.point += 1
}

func (ai *ShowdownAIPlayer) Name() string {
	return ai.name
}

func (ai *ShowdownAIPlayer) SetName(s string) {
	ai.name = s + "_AI"
}
