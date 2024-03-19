package entity

type NewMessageEvent struct {
	Sender  IMember
	Message Message
}

func (n *NewMessageEvent) GetEventName() string {
	return "NewMessageEvent"
}

func (n *NewMessageEvent) GetEventData() (IMember, Message) {
	return n.Sender, n.Message
}
