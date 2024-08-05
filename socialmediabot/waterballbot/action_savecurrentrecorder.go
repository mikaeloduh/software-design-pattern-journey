package waterballbot

import "socialmediabot/service"

func SaveCurrentRecorderAction(arg any) {
	recorder := arg.(service.TagEvent).TaggedBy.(service.IMember)

	isCurrentRecorder = func(memberId string) bool {
		return recorder.Id() == memberId
	}
}
