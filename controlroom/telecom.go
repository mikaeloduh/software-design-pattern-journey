package main

import (
	"fmt"
	"io"
	"os"
)

type Telecom struct {
	Writer io.Writer
}

func NewTelecom() *Telecom {
	return &Telecom{Writer: os.Stdout}
}

func (t *Telecom) Connect() {
	fmt.Fprint(t.Writer, "The telecom has been turned on.\n")
}

func (t *Telecom) Disconnect() {
	fmt.Fprint(t.Writer, "The telecom has been turned off.\n")
}
