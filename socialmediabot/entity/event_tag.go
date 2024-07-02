package entity

import "socialmediabot/libs"

type TagEvent struct {
	TaggedBy Taggable
	TaggedTo Taggable
	Message  Message
}

func (e TagEvent) GetData() libs.IEvent {
	return e
}
