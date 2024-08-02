package waterballbot

import (
	"socialmediabot/entity"
	"socialmediabot/libs"
)

func LogoutEventGuard(event libs.IEvent) bool {
	return event.GetData().(entity.NewLogoutEvent).OnlineCount < 10
}
