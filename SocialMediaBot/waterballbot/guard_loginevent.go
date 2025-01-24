package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

func LoginEventGuard(event libs.IEvent) bool {
	return event.GetData().(service.NewLoginEvent).OnlineCount >= 10
}
