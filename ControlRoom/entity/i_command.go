package entity

// ICommand interface
type ICommand interface {
	Execute()
	Undo()
}

// Macro
type Macro []ICommand

func (m Macro) Undo() {
	for i := len(m) - 1; i >= 0; i-- {
		m[i].Undo()
	}
}

func (m Macro) Execute() {
	for _, command := range m {
		command.Execute()
	}
}
