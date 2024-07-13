package entity

import "socialmediabot/libs"

type GoBroadcastingEvent struct {
}

func (e GoBroadcastingEvent) GetData() libs.IEvent {
	return e
}
