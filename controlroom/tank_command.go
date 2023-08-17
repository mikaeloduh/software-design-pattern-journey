package main

type MoveForwardTankCommand struct {
	tank Tank
}

func (c MoveForwardTankCommand) Execute() {
	c.tank.MoveForward()
}

func (c MoveForwardTankCommand) Undo() {
	c.tank.MoveBackward()
}
