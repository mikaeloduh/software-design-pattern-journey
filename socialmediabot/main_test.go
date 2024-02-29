package main

import (
	"bytes"
	"io"
	"socialmediabot/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_Waterball(t *testing.T) {

	t.Run("test member chat in chatroom, should print id + message", func(t *testing.T) {
		var writer bytes.Buffer

		waterball := FakeNewWaterball(&writer)
		memberId := "1"
		member := entity.NewMember(memberId)

		testMessage := "hello"
		waterball.ChatRoom.Send(*member, entity.Message{Content: testMessage})

		assert.Equal(t, memberId+": "+testMessage, writer.String())
	})
}

func FakeNewWaterball(w io.Writer) *entity.Waterball {
	return &entity.Waterball{
		Writer:   w,
		ChatRoom: entity.ChatRoom{Writer: w},
	}
}
