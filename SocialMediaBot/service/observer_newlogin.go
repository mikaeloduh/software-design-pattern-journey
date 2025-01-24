package service

import "socialmediabot/libs"

type INewLoginObserver interface {
	Update(event libs.IEvent)
}
