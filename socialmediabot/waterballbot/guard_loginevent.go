package waterballbot

import (
	"socialmediabot/entity"
	"socialmediabot/libs"
)

func LoginEventGuard(event libs.IEvent) bool {
	return event.GetData().(entity.NewLoginEvent).OnlineCount >= 10
}
