package entity

import "socialmediabot/libs"

func EnterDefaultConversationGuard(event libs.IEvent) bool {
	return event.GetData().(EnterNormalStateEvent).OnlineCount < 10
}
