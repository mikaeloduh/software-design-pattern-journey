package service

import (
	"socialmediabot/libs"
)

type NewLogoutEvent struct {
	NewLogoutMember IMember
	OnlineCount     int
}

func (e NewLogoutEvent) GetData() libs.IEvent {
	return e
}
