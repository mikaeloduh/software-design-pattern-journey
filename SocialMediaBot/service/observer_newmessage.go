package service

import "socialmediabot/libs"

type INewMessageObserver interface {
	Update(event libs.IEvent)
}
