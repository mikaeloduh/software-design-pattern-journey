package service

import "socialmediabot/libs"

type NewMessageEvent struct {
	Sender  IMember
	Message Message
}

func (e NewMessageEvent) GetData() libs.IEvent {
	return e
}
