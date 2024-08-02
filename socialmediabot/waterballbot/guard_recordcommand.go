package waterballbot

import (
	"log"
	"socialmediabot/entity"
	"socialmediabot/libs"
)

func RecordCommandGuard(event libs.IEvent) bool {
	data, ok := event.GetData().(entity.TagEvent)
	if !ok {
		log.Println("Error: Event data is not of type TagEvent")
		return false
	}

	return data.Message.Content == "record"
}
