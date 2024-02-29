package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_Waterball(t *testing.T) {

	t.Run("test member chat in chatroom", func(t *testing.T) {
		var writer bytes.Buffer
		waterball := FakeNewWaterball(&writer)
		mid := "1"
		member := NewMember(mid)

		testMessage := "hello"
		waterball.ChatRoom.Send(*member, Message{testMessage, []Tag{}})

		assert.Equal(t, mid+": "+testMessage, writer.String())
	})
}

func FakeNewWaterball(w io.Writer) *Waterball {
	return &Waterball{
		Writer:   w,
		ChatRoom: ChatRoom{Writer: w},
	}
}
