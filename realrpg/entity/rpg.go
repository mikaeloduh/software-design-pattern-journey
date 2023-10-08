package entity

type IActionHandler interface {
	Handle()
}

type RPG struct {
	units         []IUnit
	ActionHandler IActionHandler
}

func NewRPG(units []IUnit, actionHandler IActionHandler) *RPG {
	return &RPG{units: units, ActionHandler: actionHandler}
}

func (g *RPG) Run() {

}

func (g *RPG) HandleAction(target IUnit) {
	g.ActionHandler.Handle()
}
