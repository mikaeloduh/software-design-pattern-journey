package entity

import (
	"log"
	"socialmediabot/libs"
)

func StopRecordCommandGuard(event libs.IEvent) bool {
	data, ok := event.GetData().(StopRecordCommandEvent)
	if !ok {
		log.Println("Error: Event data is not of type StopRecordCommandEvent")
		return false
	}

	return data.Message.Content == "stop-recording" && data.TaggedBy == data.Recorder
}

// StopRecordCommandEvent
type StopRecordCommandEvent struct {
	TaggedBy Taggable
	TaggedTo Taggable
	Message  Message
	Recorder IMember
}

func (e StopRecordCommandEvent) GetData() libs.IEvent {
	return e
}
