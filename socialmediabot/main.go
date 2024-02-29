package main

import (
	"fmt"
	"io"
	"os"
)

type Member struct {
	Id string
}

func NewMember(id string) *Member {
	return &Member{Id: id}
}

type Tag string

type Message struct {
	Content string
	Tags    []Tag
}

type ChatRoom struct {
	Writer io.Writer
}

func (c *ChatRoom) Send(b Member, m Message) {
	_, _ = fmt.Fprint(c.Writer, fmt.Sprintf("%s: %s", b.Id, m.Content))
}

type Waterball struct {
	Writer   io.Writer
	ChatRoom ChatRoom
}

func NewWaterball() *Waterball {
	return &Waterball{
		Writer:   os.Stdout,
		ChatRoom: ChatRoom{Writer: os.Stdout},
	}
}

func main() {
	fmt.Println("hello world")
}
