package entity

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBot(t *testing.T) {
	var writer bytes.Buffer

	waterball := FakeNewWaterball(&writer)
	bot := NewBot(waterball)
	waterball.Register(bot)
	waterball.Login(bot)
	waterball.ChatRoom.Register(bot)
	member := NewMember("member_001")
	waterball.Login(member)

	t.Run("given bot is in DefaultConversationState, receiving a new message in chat should trigger a bot response", func(t *testing.T) {
		waterball.ChatRoom.Send(member, Message{Content: "hello"})

		assert.IsType(t, &DefaultConversationState{}, bot.fsm.GetState())
		assert.Equal(t, "bot_001: good to hear", getLastLine(writer.String()))
	})

	t.Run("trigger InteractingState when meet the transition rule", func(t *testing.T) {
		waterball.Login(NewMember("member_002"))
		waterball.Login(NewMember("member_003"))
		waterball.Login(NewMember("member_004"))
		waterball.Login(NewMember("member_005"))
		waterball.Login(NewMember("member_006"))
		waterball.Login(NewMember("member_007"))
		waterball.Login(NewMember("member_008"))
		waterball.Login(NewMember("member_009")) // 10th login member

		assert.Equal(t, 10, len(waterball.sessions))
		assert.IsType(t, &InteractingState{}, bot.fsm.GetState())
	})

	t.Run("given bot is in InteractingState, send new message in ChatRoom should trigger interacting-mode response", func(t *testing.T) {
		waterball.ChatRoom.Send(member, Message{Content: "hello"})

		assert.Equal(t, "bot_001: I like your idea!", getLastLine(writer.String()))
	})
}

func FakeNewWaterball(w io.Writer) *Waterball {
	return NewWaterball(w)
}

// test helper
func getLastLine(output string) string {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) == 0 {
		return ""
	}
	return lines[len(lines)-1]
}
