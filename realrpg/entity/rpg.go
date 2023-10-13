package entity

type IActionHandler interface {
	Handle()
}

type RPG struct {
	units []IUnit
}

func NewRPG(units []IUnit) *RPG {
	return &RPG{units: units}
}

func (g *RPG) Run() {
}
