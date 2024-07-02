package libs

type IGuard interface {
	Exec(event IEvent) bool
}

//type IGuard func(args ...any) bool
