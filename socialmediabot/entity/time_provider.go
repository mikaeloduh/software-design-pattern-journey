package entity

import "time"

type ITimeProvider interface {
	Sleep(d time.Duration)
}

type DefaultTimeProvider struct{}

func (DefaultTimeProvider) Sleep(d time.Duration) {
	time.Sleep(d)
}
