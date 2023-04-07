package entity

type HumanPlayer struct {
	id   int
	name string
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
