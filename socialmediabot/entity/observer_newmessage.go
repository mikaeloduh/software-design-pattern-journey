package entity

import "socialmediabot/libs"

type INewMessageObserver interface {
	Update(event libs.IEvent)
}
