package entity

// MoveForwardTankCommand
type MoveForwardTankCommand struct {
	Tank *Tank
}

func (c MoveForwardTankCommand) Execute() {
	c.Tank.MoveForward()
}

func (c MoveForwardTankCommand) Undo() {
	c.Tank.MoveBackward()
}

// MoveBackwardCommand
type MoveBackwardCommand struct {
	Tank *Tank
}

func (c MoveBackwardCommand) Execute() {
	c.Tank.MoveBackward()
}

func (c MoveBackwardCommand) Undo() {
	c.Tank.MoveForward()
}
