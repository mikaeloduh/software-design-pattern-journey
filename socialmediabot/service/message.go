package service

type Message struct {
	Sender  IMember
	Content string
	Tags    []Taggable
}

func NewMessage(sender IMember, content string, tags ...Taggable) Message {
	return Message{Sender: sender, Content: content, Tags: tags}
}
