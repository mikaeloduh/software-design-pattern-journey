package entity

type IObservable interface {
	Register(observer IObserver)
	UnRegister(observer IObserver)
	Notify()
}

type IObserver interface {
	Update(observable IObservable)
}
