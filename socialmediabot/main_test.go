package main

import (
	"bytes"
	"io"
	"socialmediabot/entity"
	"strings"
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

		assert.Equal(t, memberId+": "+testMessage, getLastLine(writer.String()))
	})

	t.Run("Tagging a member in a chat message should trigger their notification if is logged-in", func(t *testing.T) {
		var writer bytes.Buffer

		waterball := FakeNewWaterball(&writer)
		member1 := entity.NewMember("1")
		member2 := entity.NewMember("2")
		updater := &SpyUpdater{}
		member2.SetUpdater(updater)
		waterball.Login(member1)
		waterball.Login(member2)

		waterball.ChatRoom.Send(member1, entity.Message{Content: "hello", Tags: []entity.Taggable{member2}})

		assert.True(t, updater.IsCalled())
	})

	t.Run("Tagging a member in a chat message should not trigger their notification if is not logged-in", func(t *testing.T) {
		var writer bytes.Buffer

		waterball := FakeNewWaterball(&writer)
		member1 := entity.NewMember("1")
		member2 := entity.NewMember("2")
		updater := &SpyUpdater{}
		member2.SetUpdater(updater)
		waterball.Login(member1)

		waterball.ChatRoom.Send(member1, entity.Message{Content: "hello", Tags: []entity.Taggable{member2}})

		assert.False(t, updater.IsCalled())

	})
}

func FakeNewWaterball(w io.Writer) *entity.Waterball {
	return entity.NewWaterball(w)
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

// test helper
func getLastLine(output string) string {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) == 0 {
		return ""
	}
	return lines[len(lines)-1]
}
