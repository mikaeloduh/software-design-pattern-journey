package waterballbot

import (
	"log"
	"socialmediabot/libs"
	"socialmediabot/service"
)

func StopRecordCommandGuard(event libs.IEvent) bool {
	data, ok := event.GetData().(service.TagEvent)
	if !ok {
		log.Println("Error: Event data is not of type TagEvent")
		return false
	}

	return data.Message.Content == "stop-recording" && isCurrentRecorder(data.TaggedBy.Id())
}

// isCurrentRecorder
var isCurrentRecorder = func(memberId string) bool {
	return true
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
