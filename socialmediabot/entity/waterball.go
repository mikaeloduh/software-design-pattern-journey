package entity

import (
	"io"
	"os"
	"socialmediabot/libs"
)

type IChannel interface {
}

type Waterball struct {
	Writer    io.Writer
	ChatRoom  ChatRoom
	sessions  map[string]IMember
	observers []INewLoginObserver
}

func (w *Waterball) Login(member IMember) {
	w.sessions[member.Id()] = member

	w.Notify(NewLoginEvent{
		NewLoginMember: member,
		OnlineCount:    len(w.sessions),
	})
}

func (w *Waterball) Notify(event libs.IEvent) {
	for _, o := range w.observers {
		o.Update(event)
	}
}

func (w *Waterball) Register(observer INewLoginObserver) {
	w.observers = append(w.observers, observer)
}

func (w *Waterball) TagOnlineMember(event TagEvent) {
	member, exists := w.sessions[event.TaggedTo.Id()]
	if exists {
		member.Tag(event)
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
