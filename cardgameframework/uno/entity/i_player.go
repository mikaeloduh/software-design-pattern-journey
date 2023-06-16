package entity

type IPlayer interface {
	Id() int
	SetId(int)
}

type AIPlayer struct {
	id int
}

func NewAIPlayer() *AIPlayer {
	return &AIPlayer{}
}

func (p *AIPlayer) Id() int {
	return p.id
}

func (p *AIPlayer) SetId(id int) {
	p.id = id
}
