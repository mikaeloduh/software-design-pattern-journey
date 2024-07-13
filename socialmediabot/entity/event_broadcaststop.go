package entity

import "socialmediabot/libs"

type BroadcastStopEvent struct {
}

func (e BroadcastStopEvent) GetData() libs.IEvent {
	return e
}
