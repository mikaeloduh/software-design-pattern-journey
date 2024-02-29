package entity

import (
	"io"
	"os"
)

type Waterball struct {
	Writer   io.Writer
	ChatRoom ChatRoom
}

func NewWaterball() *Waterball {
	return &Waterball{
		Writer:   os.Stdout,
		ChatRoom: ChatRoom{Writer: os.Stdout},
	}
}
