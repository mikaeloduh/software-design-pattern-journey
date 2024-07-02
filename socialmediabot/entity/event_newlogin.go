package entity

import (
	"socialmediabot/libs"
)

type NewLoginEvent struct {
	NewLoginMember IMember
	OnlineCount    int
}

func (e NewLoginEvent) GetData() libs.IEvent {
	return e
}
