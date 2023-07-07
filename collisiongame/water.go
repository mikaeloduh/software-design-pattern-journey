package main

type Water struct{}

func NewWater() *Water {
	return &Water{}
}

func (w *Water) String() string {
	return "W"
}
