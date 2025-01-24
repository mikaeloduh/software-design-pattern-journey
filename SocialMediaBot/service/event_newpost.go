package service

import "socialmediabot/libs"

type NewPostEvent struct {
	PostId int
}

func (e NewPostEvent) GetData() libs.IEvent {
	return e
}
