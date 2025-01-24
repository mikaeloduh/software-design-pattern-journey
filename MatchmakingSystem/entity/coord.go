package entity

import "math"

type Coord struct {
	X float64
	Y float64
}

func (c Coord) DistanceTo(other Coord) float64 {
	return math.Sqrt(math.Pow(c.X-other.X, 2) + math.Pow(c.Y-other.Y, 2))
}
