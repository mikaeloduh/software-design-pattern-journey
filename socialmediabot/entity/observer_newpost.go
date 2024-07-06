package entity

import "socialmediabot/libs"

type INewPostObserver interface {
	Update(event libs.IEvent)
}
