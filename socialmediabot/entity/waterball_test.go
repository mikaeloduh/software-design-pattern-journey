package entity

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWaterball_ChatRoom(t *testing.T) {

	t.Run("Chatting in a chatroom should output the member's ID and message", func(t *testing.T) {
		var writer bytes.Buffer

		waterball := NewWaterball(&writer)
		memberId := "1"
		member := NewMember(memberId, USER)

		testMessage := "hello"
		waterball.ChatRoom.Send(member, NewMessage(testMessage))

		assert.Equal(t, memberId+": "+testMessage, getLastLine(writer.String()))
	})

	t.Run("Tagging a member in a chat message should trigger their notification if is logged-in", func(t *testing.T) {
		var writer bytes.Buffer

		waterball := NewWaterball(&writer)
		member1 := NewSpyMember("1")
		member2 := NewSpyMember("2")
		waterball.Login(member1)
		waterball.Login(member2)

		waterball.ChatRoom.Send(member1, NewMessage("hello", member2))

		assert.True(t, member2.IsTagCalled)
	})

	t.Run("Tagging a member in a chat message should not trigger their notification if is not logged-in", func(t *testing.T) {
		var writer bytes.Buffer

		waterball := NewWaterball(&writer)
		member1 := NewSpyMember("1")
		member2 := NewSpyMember("2")
		waterball.Login(member1) // member2 not login

		waterball.ChatRoom.Send(member1, NewMessage("hello", member2))

		assert.False(t, member2.IsTagCalled)
	})
}

type SpyMember struct {
	Member
	IsTagCalled bool
}

func NewSpyMember(id string) *SpyMember {
	return &SpyMember{
		Member:      *NewMember(id, USER),
		IsTagCalled: false,
	}
}

func (m *SpyMember) Tag(e TagEvent) {
	m.IsTagCalled = true
}

// test helper
func getLastLine(output string) string {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) == 0 {
		return ""
	}
	return lines[len(lines)-1]
}
