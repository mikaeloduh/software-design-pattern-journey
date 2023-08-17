package main

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
