package service

import "socialmediabot/libs"

type IBroadcastObserver interface {
	Update(event libs.IEvent)
}
