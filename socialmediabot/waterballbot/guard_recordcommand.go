package waterballbot

import (
	"log"
	"socialmediabot/libs"
	"socialmediabot/service"
)

func RecordCommandGuard(event libs.IEvent) bool {
	data, ok := event.GetData().(service.TagEvent)
	if !ok {
		log.Println("Error: Event data is not of type TagEvent")
		return false
	}

	return data.Message.Content == "record"
}
