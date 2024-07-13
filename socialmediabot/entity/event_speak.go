package entity

import "socialmediabot/libs"

type SpeakEvent struct {
	Speaker IMember
	Content string
}

func (e SpeakEvent) GetData() libs.IEvent {
	return e
}
