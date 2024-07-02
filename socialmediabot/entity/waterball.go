package entity

import (
	"io"
	"os"
	"socialmediabot/libs"
)

type IWaterballObserver interface {
	Update(event libs.IEvent)
}

type IChannel interface {
}

type Waterball struct {
	Writer    io.Writer
	ChatRoom  ChatRoom
	sessions  map[string]IMember
	observers []IWaterballObserver
}

func (w *Waterball) Login(member IMember) {
	w.sessions[member.Id()] = member

	for _, o := range w.observers {
		o.Update(NewLoginEvent{
			NewLoginMember: member,
			OnlineCount:    len(w.sessions),
		})
	}
}

func (w *Waterball) Register(member IWaterballObserver) {
	w.observers = append(w.observers, member)
}

func (w *Waterball) TagService(event TagEvent) {
	session, exists := w.sessions[event.TaggedTo.Id()]
	if exists {
		session.Tag(event)
	}
}

func NewWaterball(w io.Writer) *Waterball {
	waterball := &Waterball{
		Writer:   w,
		sessions: make(map[string]IMember),
	}
	waterball.ChatRoom = ChatRoom{
		Writer:    w,
		Waterball: waterball,
	}

	return waterball
}

func NewDefaultWaterball(w io.Writer) *Waterball {
	return NewWaterball(os.Stdout)
}
