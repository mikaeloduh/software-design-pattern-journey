package entity

import "socialmediabot/libs"

func LogoutEventGuard(event libs.IEvent) bool {
	return event.GetData().(NewLogoutEvent).OnlineCount < 10
}
