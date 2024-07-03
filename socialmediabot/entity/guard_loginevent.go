package entity

import "socialmediabot/libs"

func LoginEventGuard(event libs.IEvent) bool {
	return event.GetData().(NewLoginEvent).OnlineCount >= 10
}
