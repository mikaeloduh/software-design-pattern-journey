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
	tank    Tank
	telecom Telecom
}

func (c *MainController) Input(in string) {
	switch in {
	case "q":
		c.tank.MoveForward()
	}
}

func main() {
	fmt.Println("Hello world")
}
