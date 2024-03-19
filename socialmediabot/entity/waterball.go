package entity

import (
	"io"
	"os"
)

type Waterball struct {
	Writer   io.Writer
	ChatRoom ChatRoom
	sessions []IMember
}

func (w *Waterball) Login(member IMember) {
	w.sessions = append(w.sessions, member)
}

func (w *Waterball) TagService(tag Tag) {
	for _, member := range w.sessions {
		if member.Id() == tag.Id() {
			tag.Update()
		}
	}
}

func NewWaterball() *Waterball {
	waterball := &Waterball{
		Writer:   os.Stdout,
		ChatRoom: ChatRoom{Writer: os.Stdout},
	}
	waterball.ChatRoom.TagService = waterball.TagService

	return waterball
}
