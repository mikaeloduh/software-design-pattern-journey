package entity

type Message struct {
	Content string
	Tags    []Taggable
}

func NewMessage(content string, tags ...Taggable) Message {
	return Message{Content: content, Tags: tags}
}
