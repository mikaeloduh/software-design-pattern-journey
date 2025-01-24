package waterballbot

import "socialmediabot/libs"

func ClearCurrentRecorderAction(_ libs.IEvent) {
	isCurrentRecorder = func(_ string) bool {
		return true
	}
}
