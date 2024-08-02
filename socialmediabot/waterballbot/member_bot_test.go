package waterballbot

import (
	"bytes"
	"strings"
	"testing"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"

	"socialmediabot/entity"
)

func TestBot(t *testing.T) {
	var writer bytes.Buffer
	var mockClock = clock.NewMock()
	waterball := entity.NewWaterball(&writer, mockClock)
	bot := NewBot(waterball)
	waterball.Register(bot)
	waterball.Login(bot)
	waterball.ChatRoom.Register(bot)
	member001 := entity.NewMember("member_001", entity.USER)
	member002 := entity.NewMember("member_002", entity.USER)
	waterball.Login(member001)

	t.Run("given bot is in DefaultConversationState, receiving a new message in chat should trigger a bot response", func(t *testing.T) {
		waterball.ChatRoom.Send(entity.NewMessage(member001, "hello"))

		assert.IsType(t, &DefaultConversationState{}, bot.fsm.GetState())
		assert.Equal(t, "bot_001: good to hear", getLastLine(writer.String()))
	})

	t.Run("when meet the transition rule, it will trigger from DefaultConversationState to InteractingState", func(t *testing.T) {
		waterball.Login(member002)
		waterball.Login(entity.NewMember("member_003", entity.USER))
		waterball.Login(entity.NewMember("member_004", entity.USER))
		waterball.Login(entity.NewMember("member_005", entity.USER))
		waterball.Login(entity.NewMember("member_006", entity.USER))
		waterball.Login(entity.NewMember("member_007", entity.USER))
		waterball.Login(entity.NewMember("member_008", entity.USER))
		waterball.Login(entity.NewMember("member_009", entity.USER)) // 10th login member

		assert.IsType(t, &InteractingState{}, bot.fsm.GetState())
	})

	t.Run("given bot is in InteractingState, send new message in ChatRoom should trigger interacting-mode response", func(t *testing.T) {
		waterball.ChatRoom.Send(entity.NewMessage(member001, "hello"))

		assert.Equal(t, "bot_001: Hi hi", getLastLine(writer.String()))
	})

	t.Run("test state transition from NormalState to RecordState", func(t *testing.T) {
		waterball.ChatRoom.Send(entity.NewMessage(member001, "record", bot))

		assert.IsType(t, &WaitingState{}, bot.fsm.GetState())
	})

	t.Run("when a member logout and online members less than 10, nothing should happened", func(t *testing.T) {
		waterball.Logout("member_009")

		assert.IsType(t, &WaitingState{}, bot.fsm.GetState())
	})

	t.Run("Given in RecordState - Executing stop-recording command should only be successful for the recorder", func(t *testing.T) {
		waterball.ChatRoom.Send(entity.NewMessage(member002, "stop-recording", bot))

		assert.IsType(t, &WaitingState{}, bot.fsm.GetState())

		waterball.ChatRoom.Send(entity.NewMessage(member001, "stop-recording", bot))

		assert.IsType(t, &DefaultConversationState{}, bot.fsm.GetState())
	})

	t.Run("Executing king command should only be successful for admin", func(t *testing.T) {
		waterball.ChatRoom.Send(entity.NewMessage(member001, "king", bot))

		assert.IsType(t, &DefaultConversationState{}, bot.fsm.GetState())

		admin := entity.NewMember("admin_001", entity.ADMIN)
		waterball.ChatRoom.Send(entity.NewMessage(admin, "king", bot))

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
