package entity

import (
	"log"
	"socialmediabot/libs"
)

func KingCommandGuard(event libs.IEvent) bool {
	data, ok := event.GetData().(TagEvent)
	if !ok {
		log.Println("Error: Event data is not of type TagEvent")
		return false
	}

	member, ok := data.TaggedBy.(IMember)
	if !ok {
		log.Println("Error: TaggedBy is not implementing IMember")
		return false
	}

	return data.Message.Content == "king" && member.Role() == ADMIN
}
