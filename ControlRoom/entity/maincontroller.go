package entity

import (
	"controlroom/commons"
)

type MainController struct {
	commands    map[string]ICommand
	doHistory   *commons.Stack[ICommand]
	undoHistory *commons.Stack[ICommand]
}

func NewMainController() *MainController {
	return &MainController{
		commands:    make(map[string]ICommand),
		doHistory:   commons.NewStack[ICommand](),
		undoHistory: commons.NewStack[ICommand](),
	}
}

func (c *MainController) BindCommand(key string, command ICommand) {
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

func (c *MainController) Redo() {
	todo, _ := c.undoHistory.Pop()
	todo.Execute()
	c.doHistory.Push(todo)
}

func (c *MainController) Reset() {
	c.commands = make(map[string]ICommand)
}
