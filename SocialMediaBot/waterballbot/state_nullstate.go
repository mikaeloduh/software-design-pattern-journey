package waterballbot

import "socialmediabot/libs"

type NullState struct {
	libs.SuperState[IBotState]
	UnimplementedBotOperation
}
