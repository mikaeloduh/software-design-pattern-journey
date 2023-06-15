package entity

type IUnoPlayer interface {
	Id() int
	SetId(int)
}

type UnoAIPlayer struct {
	id int
}

func NewUnoAIPlayer() *UnoAIPlayer {
	return &UnoAIPlayer{}
}

func (p *UnoAIPlayer) Id() int {
	return p.id
}

func (p *UnoAIPlayer) SetId(id int) {
	p.id = id
}
