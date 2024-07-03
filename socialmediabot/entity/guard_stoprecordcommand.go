package entity

import "socialmediabot/libs"

func StopRecordCommandGuard(event libs.IEvent) bool {
	return event.GetData().(TagEvent).Message.Content == "stop-recording"
}
