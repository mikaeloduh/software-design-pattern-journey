package entity

type IEvent interface {
	GetEventName() string
}

type NewMessageEvent struct {
	evenName string
	sender   IMember
	message  Message
}

func (n *NewMessageEvent) GetEventName() string {
	return n.evenName
}

func (n *NewMessageEvent) GetEventData() (IMember, Message) {
	return n.sender, n.message
}
