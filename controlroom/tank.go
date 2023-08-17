package main

import (
	"fmt"
	"io"
	"os"
)

type Tank struct {
	Writer io.Writer
}

func NewTank() *Tank {
	return &Tank{Writer: os.Stdout}
}

func (t *Tank) MoveForward() {
	fmt.Fprint(t.Writer, "The tank has moved forward.\n")
}

func (t *Tank) MoveBackward() {
	fmt.Fprint(t.Writer, "The tank has moved backward.\n")
}
