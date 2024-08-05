package waterballbot

func ClearCurrentRecorderAction(_ any) {
	isCurrentRecorder = func(_ string) bool {
		return true
	}
}
