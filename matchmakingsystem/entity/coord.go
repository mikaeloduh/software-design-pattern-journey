package entity

import "math"

type Coord struct {
	X float64
	Y float64
}

func (c Coord) DistanceTo(other Coord) float64 {
	return math.Sqrt((c.X-other.X)*(c.X-other.X) + (c.Y-other.Y)*(c.Y-other.Y))
}
