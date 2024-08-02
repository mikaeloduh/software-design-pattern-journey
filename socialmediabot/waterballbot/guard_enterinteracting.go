package waterballbot

import "socialmediabot/libs"

func EnterInteractingGuard(event libs.IEvent) bool {
	return event.GetData().(EnterNormalStateEvent).OnlineCount >= 10
}
