package main

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"socialmediabot/entity"
	"strings"
	"testing"
)

func TestMain_Integrate(t *testing.T) {
	var writer bytes.Buffer

	waterball := entity.NewWaterball(&writer)
	bot := entity.NewBot(waterball)
	waterball.Register(bot)
	waterball.ChatRoom.Register(bot)
	waterball.Forum.Register(bot)
	waterball.Login(bot)

	member001 := entity.NewMember("member_001", entity.ADMIN)
	member004 := entity.NewMember("member_004", entity.USER)

	waterball.Login(member001)
	waterball.Login(entity.NewMember("member_002", entity.USER))
	waterball.Login(entity.NewMember("member_003", entity.USER))
	waterball.Login(member004)
	waterball.Login(entity.NewMember("member_005", entity.USER))
	waterball.Login(entity.NewMember("member_006", entity.USER))

	t.Run("2: test NormalState, DefaultConversationState", func(t *testing.T) {

		waterball.ChatRoom.Send(member001, entity.Message{Content: "Good morning, my fist day on board"})

		assert.Equal(t, "bot_001: good to hear", getLastLine(writer.String()))

		waterball.Login(entity.NewMember("member_007", entity.USER))
		waterball.ChatRoom.Send(member004, entity.Message{Content: "everyone have a good day"})

		assert.Equal(t, "bot_001: thank you", getLastLine(writer.String()))
	})

	member008 := entity.NewMember("member_008", entity.USER)

	t.Run("3: test NormalState, InteractingState", func(t *testing.T) {
		waterball.Login(member008)
		waterball.Login(entity.NewMember("member_009", entity.USER))

		waterball.ChatRoom.Send(member001, entity.Message{Content: "wow ten peoples online"})

		assert.Equal(t, "bot_001: Hi hi", getLastLine(writer.String()))

		waterball.ChatRoom.Send(member001, entity.Message{Content: "Good morning, who wants McDonald?"})

		assert.Equal(t, "bot_001: I like your idea!", getLastLine(writer.String()))
	})

	t.Run("4: Post Forum", func(t *testing.T) {
		waterball.ChatRoom.Send(member008, entity.Message{Content: "Ive post a Joke, haha"})

		assert.Equal(t, "bot_001: Hi hi", getLastLine(writer.String()))

		assert.IsType(t, &entity.InteractingState{}, bot.GetState())

		waterball.Forum.Post(member008, entity.Post{Title: "Single Responsibility Principle", Content: "Too many shit in a class, just get one shit done."})

		assert.Equal(t, "bot_001 comment in post 1: How do you guys think about it?", getLastLine(writer.String()))
	})
}

// test helper
func getLastLine(output string) string {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) == 0 {
		return ""
	}
	return lines[len(lines)-1]
}
