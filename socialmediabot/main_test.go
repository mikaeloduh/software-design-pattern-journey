package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"

	"socialmediabot/service"
	"socialmediabot/waterballbot"
)

func TestMain_Integrate(t *testing.T) {
	var writer bytes.Buffer
	var mockClock = clock.NewMock()
	waterball := service.NewWaterball(&writer, mockClock)
	bot := waterballbot.NewBot(waterball)
	waterball.Register(bot)
	waterball.ChatRoom.Register(bot)
	waterball.Forum.Register(bot)
	waterball.Broadcast.Register(bot)
	waterball.Login(bot)

	member001 := service.NewMember("member_001", service.ADMIN)
	member002 := service.NewMember("member_002", service.USER)
	member003 := service.NewMember("member_003", service.USER)
	member004 := service.NewMember("member_004", service.USER)
	member005 := service.NewMember("member_005", service.USER)
	member006 := service.NewMember("member_006", service.USER)
	member007 := service.NewMember("member_007", service.USER)
	member008 := service.NewMember("member_008", service.USER)
	member009 := service.NewMember("member_009", service.USER)

	waterball.Login(member001)
	waterball.Login(member002)
	waterball.Login(member003)
	waterball.Login(member004)
	waterball.Login(member005)
	waterball.Login(member006)

	t.Run("2: Given the bot in DefaultConversationState, chatting in ChatRoom should result in a specific reply", func(t *testing.T) {
		waterball.ChatRoom.Send(service.NewMessage(member001, "Good morning, my fist day on board"))

		assert.Equal(t, "bot_001: good to hear", getLastMessage(writer.String()))

		waterball.Login(member007)
		waterball.ChatRoom.Send(service.NewMessage(member004, "everyone have a good day", member001))

		assert.Equal(t, "bot_001: thank you", getLastMessage(writer.String()))
	})

	t.Run("3: Given the bot in InteractingState, chatting in ChatRoom should result in a specific reply", func(t *testing.T) {
		waterball.Login(member008)
		waterball.Login(member009)

		waterball.ChatRoom.Send(service.NewMessage(member001, "wow ten peoples online"))

		assert.Equal(t, "bot_001: Hi hi", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(service.NewMessage(member001, "Good morning, who wants McDonald?"))

		assert.Equal(t, "bot_001: I like your idea!", getLastMessage(writer.String()))
	})

	t.Run("4: Given the bot in InteractingState, Posting in Forum should result in a specific reply", func(t *testing.T) {
		waterball.ChatRoom.Send(service.NewMessage(member008, "Ive post a Joke, haha"))

		assert.Equal(t, "bot_001: Hi hi", getLastMessage(writer.String()))

		assert.IsType(t, &waterballbot.InteractingState{}, bot.GetState())

		waterball.Forum.Post(member008, service.Post{Title: "Single Responsibility Principle", Content: "Too many shit in a class, just get one shit done."})

		assert.Equal(t, "bot_001 comment in post 1: How do you guys think about it?", getLastMessage(writer.String()))
	})

	t.Run("5: Given the bot in NormalState, when making the 'king' command, the bot should transition to QuestioningState", func(t *testing.T) {
		assert.IsType(t, &waterballbot.InteractingState{}, bot.GetState())

		waterball.ChatRoom.Send(service.NewMessage(member001, "king", bot))

		assert.Equal(t, "bot_001: I like your idea!", getNthLastMessage(writer.String(), 3))
		assert.Equal(t, "bot_001: KnowledgeKing is started!", getNthLastMessage(writer.String(), 2))
		assert.IsType(t, &waterballbot.QuestioningState{}, bot.GetState())
	})

	t.Run("5-Q1: When QuestioningState is initialized, it should begin with the first question ", func(t *testing.T) {
		assert.Equal(t, "bot_001: 請問哪個 SQL 語句用於選擇所有的行？\nA) SELECT *\nB) SELECT ALL\nC) SELECT ROWS\nD) SELECT DATA", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(service.NewMessage(member006, "A", bot))

		assert.Equal(t, "bot_001: Congrats! you got the answer!", getNthLastMessage(writer.String(), 2))
	})

	t.Run("5-Q2: Following the previous test, when the first question has been answered, the second question should be presented", func(t *testing.T) {
		assert.Equal(t, "bot_001: 請問哪個 CSS 屬性可用於設置文字的顏色？\nA) text-align\nB) font-size\nC) color\nD) padding", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(service.NewMessage(member008, "C", bot))

		assert.Equal(t, "bot_001: Congrats! you got the answer!", getNthLastMessage(writer.String(), 2))
	})

	t.Run("5-Q3: Only the correct answer will be counted ", func(t *testing.T) {
		assert.Equal(t, "bot_001: 請問在計算機科學中，「XML」代表什麼？\nA) Extensible Markup Language\nB) Extensible Modeling Language\nC) Extended Markup Language\nD) Extended Modeling Language", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(service.NewMessage(member003, "C", bot))
		waterball.ChatRoom.Send(service.NewMessage(member008, "A", bot))

		assert.Equal(t, "bot_001: Congrats! you got the answer!", getNthLastMessage(writer.String(), 4))
		assert.IsType(t, &waterballbot.ThanksForJoiningState{}, bot.GetState())
	})

	t.Run("5-end: ThanksForJoiningState should release the game result and wait for 5 seconds before transiting to InteractingState", func(t *testing.T) {
		assert.Equal(t, "bot_001 go broadcasting...", getNthLastMessage(writer.String(), 3))
		assert.Equal(t, "bot_001 speaking: The winner is member_008", getNthLastMessage(writer.String(), 2))

		mockClock.Add(5 * time.Second)

		assert.IsType(t, &waterballbot.InteractingState{}, bot.GetState())
	})

	t.Run("6-1: Given the bot is in WaitingState, when the broadcast starts, the bot should transition to RecordingState", func(t *testing.T) {
		waterball.ChatRoom.Send(service.NewMessage(member003, "record", bot))
		assert.IsType(t, &waterballbot.WaitingState{}, bot.GetState())

		assert.Equal(t, "bot_001: Hi hi", getLastMessage(writer.String()))

		err := waterball.Broadcast.GoBroadcasting(member004)
		assert.NoError(t, err)

		assert.Equal(t, "member_004 go broadcasting...", getLastMessage(writer.String()))
		assert.IsType(t, &waterballbot.RecordingState{}, bot.GetState())

		waterball.Broadcast.Transmit(service.NewSpeak(member004, "Good morning guys"))
		assert.Equal(t, "member_004 speaking: Good morning guys", getLastMessage(writer.String()))

		waterball.Broadcast.Transmit(service.NewSpeak(member004, "Have you had breakfast yet?"))
		assert.Equal(t, "member_004 speaking: Have you had breakfast yet?", getLastMessage(writer.String()))
	})

	t.Run("6-2: When the broadcast stops, the bot should output the recorded text", func(t *testing.T) {
		err := waterball.Broadcast.StopBroadcasting(member004)
		assert.NoError(t, err)
		assert.Equal(t, "member_004 stop broadcasting", getNthLastMessage(writer.String(), 2))

		assert.Equal(t, "bot_001: [Record Replay] Good morning guys\nHave you had breakfast yet?", getLastMessage(writer.String()))

		waterball.ChatRoom.Send(service.NewMessage(member003, "stop-recording", bot))

		assert.IsType(t, &waterballbot.InteractingState{}, bot.GetState())
	})

	t.Run("7: Given the bot in InteractingState, when a user logout and online count is less than 10, it should transition to DefaultConversationState", func(t *testing.T) {
		waterball.Logout(member009.Id())
		waterball.Logout(member008.Id())
		waterball.Logout(member007.Id())
		waterball.Logout(member006.Id())

		waterball.ChatRoom.Send(service.NewMessage(member001, "Yeah, people went off-line"))

		assert.IsType(t, &waterballbot.DefaultConversationState{}, bot.GetState())
		assert.Equal(t, "bot_001: How are you", getLastMessage(writer.String()))
	})
}

func TestMain_Integrate_Timeout(t *testing.T) {
	var writer bytes.Buffer
	var mockClock = clock.NewMock()
	waterball := service.NewWaterball(&writer, mockClock)
	bot := waterballbot.NewBot(waterball)
	waterball.Register(bot)
	waterball.ChatRoom.Register(bot)
	waterball.Forum.Register(bot)
	waterball.Broadcast.Register(bot)
	waterball.Login(bot)

	member001 := service.NewMember("member_001", service.ADMIN)
	waterball.Login(member001)

	t.Run("Given KnowledgeKingState is initialized to QuestioningState - ", func(t *testing.T) {
		waterball.ChatRoom.Send(service.NewMessage(member001, "king", bot))

		assert.IsType(t, &waterballbot.QuestioningState{}, bot.GetState())
	})

	t.Run("After an hour of inactivity - QuestioningState should end automatically", func(t *testing.T) {
		mockClock.Add(1 * time.Hour)

		assert.IsType(t, &waterballbot.ThanksForJoiningState{}, bot.GetState())
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
