package main

import (
	"bytes"
	"fmt"
)

type Telecom struct {
	Writer *bytes.Buffer
}

func (t *Telecom) Connect() {
	fmt.Fprint(t.Writer, "The telecom has been turned on.\n")
}

func (t *Telecom) Disconnect() {
	fmt.Fprint(t.Writer, "The telecom has been turned off.\n")
}
