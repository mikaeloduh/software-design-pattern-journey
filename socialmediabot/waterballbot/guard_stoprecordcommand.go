package waterballbot

import (
	"log"
	"socialmediabot/libs"
	"socialmediabot/service"
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
	TaggedBy service.Taggable
	TaggedTo service.Taggable
	Message  service.Message
	Recorder service.IMember
}

func (e StopRecordCommandEvent) GetData() libs.IEvent {
	return e
}
