package libs

func PositiveTestGuard(_ IEvent) bool {
	return true
}

func NegativeTestGuard(_ IEvent) bool {
	return false
}

func NoAction(_ any) {}
