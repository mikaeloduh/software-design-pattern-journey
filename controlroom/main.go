package main

import (
	"bytes"
	"fmt"
	"io"
)

type Tank struct {
	Writer io.Writer
}

func (t *Tank) MoveForward() {
	fmt.Fprint(t.Writer, "The tank has moved forward.\n")
}

func (t *Tank) MoveBackward() {
	fmt.Fprint(t.Writer, "The tank has moved backward.\n")
}

type Telecom struct {
	Writer *bytes.Buffer
}

func (t *Telecom) Connect() {
	fmt.Fprint(t.Writer, "The telecom has been turned on.\n")
}

func (t *Telecom) Disconnect() {
	fmt.Fprint(t.Writer, "The telecom has been turned off.\n")
}

type MainController struct {
	commands    map[string]ICommand
	doHistory   *Stack[ICommand]
	undoHistory *Stack[ICommand]
}

func NewMainController() *MainController {
	return &MainController{
		commands:    make(map[string]ICommand),
		doHistory:   NewStack[ICommand](),
		undoHistory: NewStack[ICommand](),
	}
}

func (c *MainController) SetCommand(key string, command ICommand) {
	c.commands[key] = command
}

func (c *MainController) Input(in string) {
	todo := c.commands[in]
	todo.Execute()
	c.doHistory.Push(todo)
	c.undoHistory.Clear()
}

func (c *MainController) Undo() {
	todo, _ := c.doHistory.Pop()
	todo.Undo()
	c.undoHistory.Push(todo)
}

type ICommand interface {
	Execute()
	Undo()
}

type MoveForwardTankCommand struct {
	tank Tank
}

func (c MoveForwardTankCommand) Execute() {
	c.tank.MoveForward()
}

func (c MoveForwardTankCommand) Undo() {
	c.tank.MoveBackward()
}

func main() {
	fmt.Println("Hello world")
}
