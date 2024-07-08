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
	member002 := entity.NewMember("member_002", entity.USER)
	member003 := entity.NewMember("member_003", entity.USER)
	member004 := entity.NewMember("member_004", entity.USER)
	member006 := entity.NewMember("member_006", entity.USER)

	waterball.Login(member001)
	waterball.Login(member002)
	waterball.Login(member003)
	waterball.Login(member004)
	waterball.Login(entity.NewMember("member_005", entity.USER))
	waterball.Login(member006)

	t.Run("2: test NormalState, DefaultConversationState", func(t *testing.T) {
		waterball.ChatRoom.Send(member001, entity.Message{Content: "Good morning, my fist day on board"})

		assert.Equal(t, "bot_001: good to hear", getLastMessage(writer.String()))

		waterball.Login(entity.NewMember("member_007", entity.USER))
		waterball.ChatRoom.Send(member004, entity.NewMessage("everyone have a good day", member001))

		assert.Equal(t, "bot_001: thank you", getLastMessage(writer.String()))
	})

	member008 := entity.NewMember("member_008", entity.USER)

	t.Run("3: test NormalState, InteractingState", func(t *testing.T) {
		waterball.Login(member008)
		waterball.Login(entity.NewMember("member_009", entity.USER))

		waterball.ChatRoom.Send(member001, entity.NewMessage("wow ten peoples online"))

		assert.Equal(t, "bot_001: Hi hi", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(member001, entity.NewMessage("Good morning, who wants McDonald?"))

		assert.Equal(t, "bot_001: I like your idea!", getLastMessage(writer.String()))
	})

	t.Run("4: Post Forum", func(t *testing.T) {
		waterball.ChatRoom.Send(member008, entity.NewMessage("Ive post a Joke, haha"))

		assert.Equal(t, "bot_001: Hi hi", getLastMessage(writer.String()))

		assert.IsType(t, &entity.InteractingState{}, bot.GetState())

		waterball.Forum.Post(member008, entity.Post{Title: "Single Responsibility Principle", Content: "Too many shit in a class, just get one shit done."})

		assert.Equal(t, "bot_001 comment in post 1: How do you guys think about it?", getLastMessage(writer.String()))
	})

	t.Run("5: test KnowledgeKingState", func(t *testing.T) {
		assert.IsType(t, &entity.InteractingState{}, bot.GetState())

		waterball.ChatRoom.Send(member001, entity.NewMessage("king", bot))

		assert.Equal(t, "bot_001: I like your idea!", getNthLastMessage(writer.String(), 3))
		assert.IsType(t, &entity.QuestioningState{}, bot.GetState())
		assert.Equal(t, "bot_001: KnowledgeKing is started!", getNthLastMessage(writer.String(), 2))
	})

	t.Run("5-Q1: starting QuestioningState should begin with the first question ", func(t *testing.T) {
		assert.Equal(t, "bot_001: 請問哪個 SQL 語句用於選擇所有的行？\nA) SELECT *\nB) SELECT ALL\nC) SELECT ROWS\nD) SELECT DATA", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(member006, entity.NewMessage("A", bot))

		assert.Equal(t, "bot_001: Congrats! you got the answer!", getNthLastMessage(writer.String(), 2))
	})

	t.Run("5-Q2: questions should be asked sequentially in order", func(t *testing.T) {
		assert.Equal(t, "bot_001: 請問哪個 CSS 屬性可用於設置文字的顏色？\nA) text-align\nB) font-size\nC) color\nD) padding", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(member008, entity.NewMessage("C", bot))

		assert.Equal(t, "bot_001: Congrats! you got the answer!", getNthLastMessage(writer.String(), 2))
	})

	t.Run("5-Q3: submitting the correct answer should count only the right answer", func(t *testing.T) {
		assert.Equal(t, "bot_001: 請問在計算機科學中，「XML」代表什麼？\nA) Extensible Markup Language\nB) Extensible Modeling Language\nC) Extended Markup Language\nD) Extended Modeling Language", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(member003, entity.NewMessage("C", bot))
		waterball.ChatRoom.Send(member002, entity.NewMessage("A", bot))

		assert.Equal(t, "bot_001: Congrats! you got the answer!", getNthLastMessage(writer.String(), 3))
	})

	t.Run("5-end: exiting QuestioningState should enter ThanksForJoiningState", func(t *testing.T) {
		assert.IsType(t, &entity.ThanksForJoiningState{}, bot.GetState())

		assert.Equal(t, "bot_001 go broadcasting...", getNthLastMessage(writer.String(), 2))
		assert.Equal(t, "bot_001 speaking: The winner is member_008", getLastMessage(writer.String()))
	})
}

// test helper
func getLastMessage(output string) string {
	return getNthLastMessage(output, 1)
}

func getNthLastMessage(output string, nthLast int) string {
	lines := strings.Split(strings.TrimSpace(output), "\f")
	if len(lines) == 0 {
		return ""
	}
	return lines[len(lines)-nthLast]
}
