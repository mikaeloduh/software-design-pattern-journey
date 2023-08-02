package main

import (
	"fmt"
	"io"
	"os"
)

// Channel type
type Channel struct {
	Name        string
	Videos      []*Video
	Subscribers []*Subscriber
	Writer      io.Writer
}

func NewChannel(name string) *Channel {
	return &Channel{Name: name, Writer: os.Stdout}
}

func (c *Channel) Subscribe(s *Subscriber) {
	c.Subscribers = append(c.Subscribers, s)

	fmt.Fprintf(c.Writer, "%s 訂閱了 %s。", s.Name, c.Name)
}

func (c *Channel) Unsubscribe(s *Subscriber) {
	var temp []*Subscriber
	for _, v := range c.Subscribers {
		if v != s {
			temp = append(temp, v)
		}
	}
	c.Subscribers = temp

	fmt.Fprintf(c.Writer, "%s 解除訂閱了 %s。", s.Name, c.Name)
}

func (c *Channel) Upload(v *Video) {
	c.Videos = append(c.Videos, v)
	c.Notify(v)
}

func (c *Channel) Notify(v *Video) {
	fmt.Fprintf(c.Writer, "頻道 %s 上架了一則新影片 \"%s\"。", c.Name, v.Title)

	for _, s := range c.Subscribers {
		s.OnNotify(c, v)
	}
}
