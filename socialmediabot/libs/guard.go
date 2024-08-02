package libs

type Guard func(event IEvent) bool

type IGuard interface {
	Do(event IEvent) bool
}
