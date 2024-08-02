package service

import "socialmediabot/libs"

// IBroadcastObserver
type IBroadcastObserver interface {
	Update(event libs.IEvent)
}
