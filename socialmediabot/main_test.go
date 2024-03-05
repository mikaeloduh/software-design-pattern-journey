package main

import (
	"bytes"
	"io"
	"socialmediabot/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain_Waterball(t *testing.T) {

	t.Run("Chatting in a chatroom should output the member's ID and message", func(t *testing.T) {
		var writer bytes.Buffer

		waterball := FakeNewWaterball(&writer)
		memberId := "1"
		member := entity.NewMember(memberId)

		testMessage := "hello"
		waterball.ChatRoom.Send(member, entity.Message{Content: testMessage})

		assert.Equal(t, memberId+": "+testMessage, writer.String())
	})

	t.Run("Tagging a member in a chat message should trigger their notification", func(t *testing.T) {
		var writer bytes.Buffer

		waterball := FakeNewWaterball(&writer)
		member1 := entity.NewMember("1")
		member2 := entity.NewMember("2")
		updater := &SpyUpdater{}
		member2.SetUpdater(updater)

		waterball.ChatRoom.Send(member1, entity.Message{Content: "hello", Tags: []entity.Tag{member2}})

		assert.True(t, updater.IsCalled())
	})
}

func FakeNewWaterball(w io.Writer) *entity.Waterball {
	return &entity.Waterball{
		Writer:   w,
		ChatRoom: entity.ChatRoom{Writer: w},
	}
}

type SpyUpdater struct {
	isCalled bool
}

func (b *SpyUpdater) Do() {
	b.isCalled = true
}

func (b *SpyUpdater) IsCalled() bool {
	return b.isCalled
}
