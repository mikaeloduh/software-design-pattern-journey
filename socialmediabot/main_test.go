package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"

	"socialmediabot/entity"
)

func TestMain_Integrate(t *testing.T) {
	var writer bytes.Buffer
	var mockClock = clock.NewMock()
	waterball := entity.NewWaterball(&writer, mockClock)
	bot := entity.NewBot(waterball)
	waterball.Register(bot)
	waterball.ChatRoom.Register(bot)
	waterball.Forum.Register(bot)
	waterball.Broadcast.Register(bot)
	waterball.Login(bot)

	member001 := entity.NewMember("member_001", entity.ADMIN)
	member002 := entity.NewMember("member_002", entity.USER)
	member003 := entity.NewMember("member_003", entity.USER)
	member004 := entity.NewMember("member_004", entity.USER)
	member005 := entity.NewMember("member_005", entity.USER)
	member006 := entity.NewMember("member_006", entity.USER)
	member007 := entity.NewMember("member_007", entity.USER)
	member008 := entity.NewMember("member_008", entity.USER)
	member009 := entity.NewMember("member_009", entity.USER)

	waterball.Login(member001)
	waterball.Login(member002)
	waterball.Login(member003)
	waterball.Login(member004)
	waterball.Login(member005)
	waterball.Login(member006)

	t.Run("2: test NormalState, DefaultConversationState", func(t *testing.T) {
		waterball.ChatRoom.Send(entity.NewMessage(member001, "Good morning, my fist day on board"))

		assert.Equal(t, "bot_001: good to hear", getLastMessage(writer.String()))

		waterball.Login(member007)
		waterball.ChatRoom.Send(entity.NewMessage(member004, "everyone have a good day", member001))

		assert.Equal(t, "bot_001: thank you", getLastMessage(writer.String()))
	})

	t.Run("3: test NormalState, InteractingState", func(t *testing.T) {
		waterball.Login(member008)
		waterball.Login(member009)

		waterball.ChatRoom.Send(entity.NewMessage(member001, "wow ten peoples online"))

		assert.Equal(t, "bot_001: Hi hi", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(entity.NewMessage(member001, "Good morning, who wants McDonald?"))

		assert.Equal(t, "bot_001: I like your idea!", getLastMessage(writer.String()))
	})

	t.Run("4: Post Forum", func(t *testing.T) {
		waterball.ChatRoom.Send(entity.NewMessage(member008, "Ive post a Joke, haha"))

		assert.Equal(t, "bot_001: Hi hi", getLastMessage(writer.String()))

		assert.IsType(t, &entity.InteractingState{}, bot.GetState())

		waterball.Forum.Post(member008, entity.Post{Title: "Single Responsibility Principle", Content: "Too many shit in a class, just get one shit done."})

		assert.Equal(t, "bot_001 comment in post 1: How do you guys think about it?", getLastMessage(writer.String()))
	})

	t.Run("5: test KnowledgeKingState", func(t *testing.T) {
		assert.IsType(t, &entity.InteractingState{}, bot.GetState())

		waterball.ChatRoom.Send(entity.NewMessage(member001, "king", bot))

		assert.Equal(t, "bot_001: I like your idea!", getNthLastMessage(writer.String(), 3))
		assert.Equal(t, "bot_001: KnowledgeKing is started!", getNthLastMessage(writer.String(), 2))
		assert.IsType(t, &entity.QuestioningState{}, bot.GetState())
	})

	t.Run("5-Q1: starting QuestioningState should begin with the first question ", func(t *testing.T) {
		assert.Equal(t, "bot_001: 請問哪個 SQL 語句用於選擇所有的行？\nA) SELECT *\nB) SELECT ALL\nC) SELECT ROWS\nD) SELECT DATA", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(entity.NewMessage(member006, "A", bot))

		assert.Equal(t, "bot_001: Congrats! you got the answer!", getNthLastMessage(writer.String(), 2))
	})

	t.Run("5-Q2: questions should be asked sequentially in order", func(t *testing.T) {
		assert.Equal(t, "bot_001: 請問哪個 CSS 屬性可用於設置文字的顏色？\nA) text-align\nB) font-size\nC) color\nD) padding", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(entity.NewMessage(member008, "C", bot))

		assert.Equal(t, "bot_001: Congrats! you got the answer!", getNthLastMessage(writer.String(), 2))
	})

	t.Run("5-Q3: submitting the correct answer should count only the right answer", func(t *testing.T) {
		assert.Equal(t, "bot_001: 請問在計算機科學中，「XML」代表什麼？\nA) Extensible Markup Language\nB) Extensible Modeling Language\nC) Extended Markup Language\nD) Extended Modeling Language", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(entity.NewMessage(member003, "C", bot))
		waterball.ChatRoom.Send(entity.NewMessage(member008, "A", bot))

		assert.Equal(t, "bot_001: Congrats! you got the answer!", getNthLastMessage(writer.String(), 4))
	})

	t.Run("5-end: exiting QuestioningState should enter ThanksForJoiningState", func(t *testing.T) {
		assert.IsType(t, &entity.ThanksForJoiningState{}, bot.GetState())
		assert.Equal(t, "bot_001 go broadcasting...", getNthLastMessage(writer.String(), 3))
		assert.Equal(t, "bot_001 speaking: The winner is member_008", getNthLastMessage(writer.String(), 2))

		mockClock.Add(5 * time.Second)

		assert.IsType(t, &entity.InteractingState{}, bot.GetState())
	})

	t.Run("6-1: while in WaitingState, stating a broadcast should enter RecordingState", func(t *testing.T) {
		waterball.ChatRoom.Send(entity.NewMessage(member003, "record", bot))
		assert.IsType(t, &entity.WaitingState{}, bot.GetState())

		assert.Equal(t, "bot_001: Hi hi", getLastMessage(writer.String()))

		err := waterball.Broadcast.GoBroadcasting(member004)
		assert.NoError(t, err)

		assert.Equal(t, "member_004 go broadcasting...", getLastMessage(writer.String()))
		assert.IsType(t, &entity.RecordingState{}, bot.GetState())

		waterball.Broadcast.Transmit(entity.NewSpeak(member004, "Good morning guys"))
		assert.Equal(t, "member_004 speaking: Good morning guys", getLastMessage(writer.String()))

		waterball.Broadcast.Transmit(entity.NewSpeak(member004, "Have you had breakfast yet?"))
		assert.Equal(t, "member_004 speaking: Have you had breakfast yet?", getLastMessage(writer.String()))
	})

	t.Run("6-2: when broadcast stop, bot should output the recorded text", func(t *testing.T) {
		err := waterball.Broadcast.StopBroadcasting(member004)
		assert.NoError(t, err)
		assert.Equal(t, "member_004 stop broadcasting", getNthLastMessage(writer.String(), 2))

		assert.Equal(t, "bot_001: [Record Replay] Good morning guys\nHave you had breakfast yet?", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(entity.NewMessage(member003, "stop-recording", bot))

		assert.IsType(t, &entity.InteractingState{}, bot.GetState())
	})

	t.Run("7: when online users are less than 10, should transit to DefaultConversationState", func(t *testing.T) {
		waterball.Logout(member009.Id())
		waterball.Logout(member008.Id())
		waterball.Logout(member007.Id())
		waterball.Logout(member006.Id())

		waterball.ChatRoom.Send(entity.NewMessage(member001, "Yeah, people went off-line"))

		assert.IsType(t, &entity.DefaultConversationState{}, bot.GetState())
		assert.Equal(t, "bot_001: How are you", getLastMessage(writer.String()))
	})
}

func TestMain_Integrate_Timeout(t *testing.T) {
	var writer bytes.Buffer
	var mockClock = clock.NewMock()
	waterball := entity.NewWaterball(&writer, mockClock)
	bot := entity.NewBot(waterball)
	waterball.Register(bot)
	waterball.ChatRoom.Register(bot)
	waterball.Forum.Register(bot)
	waterball.Broadcast.Register(bot)
	waterball.Login(bot)

	member001 := entity.NewMember("member_001", entity.ADMIN)
	waterball.Login(member001)

	t.Run("Given KnowledgeKingState is initialized to QuestioningState - ", func(t *testing.T) {
		waterball.ChatRoom.Send(entity.NewMessage(member001, "king", bot))

		assert.IsType(t, &entity.QuestioningState{}, bot.GetState())
	})

	t.Run("When an hour of inactivity in QuestioningState - QuestioningState should end automatically", func(t *testing.T) {
		mockClock.Add(1 * time.Hour)

		assert.IsType(t, &entity.ThanksForJoiningState{}, bot.GetState())
		assert.Equal(t, "bot_001 go broadcasting...", getNthLastMessage(writer.String(), 3))
		assert.Equal(t, "bot_001 speaking: Tie!", getNthLastMessage(writer.String(), 2))
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

func printMessage(output string) {
	lines := strings.Split(strings.TrimSpace(output), "\f")
	if len(lines) == 0 {
		fmt.Println("")
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
