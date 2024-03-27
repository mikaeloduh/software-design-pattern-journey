package entity

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBot(t *testing.T) {
	t.Run("given bot is in Default Conversation state, receiving a new message in chat should trigger a bot response", func(t *testing.T) {
		var writer bytes.Buffer

		waterball := FakeNewWaterball(&writer)
		bot := NewBot(waterball)
		waterball.Login(bot)
		waterball.ChatRoom.Register(bot)

		member := NewMember("1")
		waterball.Login(member)
		waterball.ChatRoom.Send(member, Message{Content: "hello"})

		botId := "bot_001"
		expectedMessage := "good to hear"
		assert.Equal(t, botId+": "+expectedMessage, getLastLine(writer.String()))
	})
}

func FakeNewWaterball(w io.Writer) Waterball {
	waterball := Waterball{
		Writer:   w,
		ChatRoom: ChatRoom{Writer: w},
	}
	waterball.ChatRoom.TagService = waterball.TagService

	return waterball
}

// test helper
func getLastLine(output string) string {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) == 0 {
		return ""
	}
	return lines[len(lines)-1]
}
