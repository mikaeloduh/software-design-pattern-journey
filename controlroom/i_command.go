package main

type ICommand interface {
	Execute()
	Undo()
}
