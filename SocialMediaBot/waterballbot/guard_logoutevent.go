package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

func LogoutEventGuard(event libs.IEvent) bool {
	return event.GetData().(service.NewLogoutEvent).OnlineCount < 10
}
