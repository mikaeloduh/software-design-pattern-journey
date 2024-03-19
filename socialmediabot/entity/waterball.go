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

func (w *Waterball) TagService(event TagEvent) {
	for _, member := range w.sessions {
		if member.Id() == event.TaggedTo.Id() {
			member.Tag(event)
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
