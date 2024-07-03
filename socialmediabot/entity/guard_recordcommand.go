package entity

import "socialmediabot/libs"

func RecordCommandGuard(event libs.IEvent) bool {
	return event.GetData().(TagEvent).Message.Content == "record"
}
