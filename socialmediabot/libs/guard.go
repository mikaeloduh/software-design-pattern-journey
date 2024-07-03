package libs

//type IGuard interface {
//	Exec(event IEvent) bool
//}

type Guard func(event IEvent) bool
