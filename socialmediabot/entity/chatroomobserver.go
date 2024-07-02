package entity

import "socialmediabot/libs"

type IChatRoomObserver interface {
	Update(event libs.IEvent)
}
