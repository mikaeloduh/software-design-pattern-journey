package entity

import (
	"log"
	"socialmediabot/libs"
)

func RecordCommandGuard(event libs.IEvent) bool {
	data, ok := event.GetData().(TagEvent)
	if !ok {
		log.Println("Error: Event data is not of type TagEvent")
		return false
	}

	return data.Message.Content == "record"
}
