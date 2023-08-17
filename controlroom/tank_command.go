package main

// MoveForwardTankCommand
type MoveForwardTankCommand struct {
	tank *Tank
}

func (c MoveForwardTankCommand) Execute() {
	c.tank.MoveForward()
}

func (c MoveForwardTankCommand) Undo() {
	c.tank.MoveBackward()
}

// MoveBackwardCommand
type MoveBackwardCommand struct {
	tank *Tank
}

func (c MoveBackwardCommand) Execute() {
	c.tank.MoveBackward()
}

func (c MoveBackwardCommand) Undo() {
	c.tank.MoveForward()
}
