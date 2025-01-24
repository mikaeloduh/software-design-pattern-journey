package entity

type Obstacle struct {
	Position *Position
}

func (o *Obstacle) SetPosition(p *Position) {
	o.Position = p

}

func (o *Obstacle) GetPosition() *Position {
	return o.Position
}
