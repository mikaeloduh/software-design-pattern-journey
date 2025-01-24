package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

func SaveCurrentRecorderAction(event libs.IEvent) {
	recorder := event.(service.TagEvent).TaggedBy.(service.IMember)

	isCurrentRecorder = func(memberId string) bool {
		return recorder.Id() == memberId
	}
}
