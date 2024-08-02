package waterballbot

import (
	"bytes"
	"strings"
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"

	"socialmediabot/service"
)

func TestBot(t *testing.T) {
	var writer bytes.Buffer
	var mockClock = clock.NewMock()
	waterball := service.NewWaterball(&writer, mockClock)
	bot := NewBot(waterball)
	waterball.Register(bot)
	waterball.Login(bot)
	waterball.ChatRoom.Register(bot)
	member001 := service.NewMember("member_001", service.USER)
	member002 := service.NewMember("member_002", service.USER)
	waterball.Login(member001)

	t.Run("given bot is in DefaultConversationState, receiving a new message in chat should trigger a bot response", func(t *testing.T) {
		waterball.ChatRoom.Send(service.NewMessage(member001, "hello"))

		assert.IsType(t, &DefaultConversationState{}, bot.fsm.GetState())
		assert.Equal(t, "bot_001: good to hear", getLastLine(writer.String()))
	})

	t.Run("when meet the transition rule, it will trigger from DefaultConversationState to InteractingState", func(t *testing.T) {
		waterball.Login(member002)
		waterball.Login(service.NewMember("member_003", service.USER))
		waterball.Login(service.NewMember("member_004", service.USER))
		waterball.Login(service.NewMember("member_005", service.USER))
		waterball.Login(service.NewMember("member_006", service.USER))
		waterball.Login(service.NewMember("member_007", service.USER))
		waterball.Login(service.NewMember("member_008", service.USER))
		waterball.Login(service.NewMember("member_009", service.USER)) // 10th login member

		assert.IsType(t, &InteractingState{}, bot.fsm.GetState())
	})

	t.Run("given bot is in InteractingState, send new message in ChatRoom should trigger interacting-mode response", func(t *testing.T) {
		waterball.ChatRoom.Send(service.NewMessage(member001, "hello"))

		assert.Equal(t, "bot_001: Hi hi", getLastLine(writer.String()))
	})

	t.Run("test state transition from NormalState to RecordState", func(t *testing.T) {
		waterball.ChatRoom.Send(service.NewMessage(member001, "record", bot))

		assert.IsType(t, &WaitingState{}, bot.fsm.GetState())
	})

	t.Run("when a member logout and online members less than 10, nothing should happened", func(t *testing.T) {
		waterball.Logout("member_009")

		assert.IsType(t, &WaitingState{}, bot.fsm.GetState())
	})

	t.Run("Given in RecordState - Executing stop-recording command should only be successful for the recorder", func(t *testing.T) {
		waterball.ChatRoom.Send(service.NewMessage(member002, "stop-recording", bot))

		assert.IsType(t, &WaitingState{}, bot.fsm.GetState())

		waterball.ChatRoom.Send(service.NewMessage(member001, "stop-recording", bot))

		assert.IsType(t, &DefaultConversationState{}, bot.fsm.GetState())
	})

	t.Run("Executing king command should only be successful for admin", func(t *testing.T) {
		waterball.ChatRoom.Send(service.NewMessage(member001, "king", bot))

		assert.IsType(t, &DefaultConversationState{}, bot.fsm.GetState())

		admin := service.NewMember("admin_001", service.ADMIN)
		waterball.ChatRoom.Send(service.NewMessage(admin, "king", bot))

		assert.IsType(t, &QuestioningState{}, bot.fsm.GetState())
	})
}

// test helper
func getLastLine(output string) string {
	lines := strings.Split(strings.TrimSpace(output), "\f")
	if len(lines) == 0 {
		return ""
	}
	return lines[len(lines)-1]
}
