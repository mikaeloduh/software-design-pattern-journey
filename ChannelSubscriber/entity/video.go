package entity

import (
	"fmt"
	"io"
)

// Video type
type Video struct {
	Title       string
	Description string
	Length      int
	LikeBy      []*Subscriber
	Writer      io.Writer
}

func (v *Video) Like(s *Subscriber) {
	v.LikeBy = append(v.LikeBy, s)

	fmt.Fprintf(v.Writer, "%s 對影片 \"%s\" 按讚。", s.Name, v.Title)
}
