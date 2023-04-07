package entity

type IPlayer interface {
}

type Player struct {
	id   int
	name string
}

func NewPlayer(id int) *Player {
	return &Player{id: id}
}

func (p *Player) Id() int {
	return p.id
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) SetName(name string) {
	p.name = name
}
