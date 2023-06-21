package entity

import "fmt"

type Color string

const (
	Red    Color = "Red"
	Blue   Color = "Blue"
	Green  Color = "Green"
	Yellow Color = "Yellow"
)

type Value int

const (
	Zero  Value = 0
	One   Value = 1
	Two   Value = 2
	Three Value = 3
	Four  Value = 4
	Five  Value = 5
	Six   Value = 6
	Seven Value = 7
	Eight Value = 8
	Nine  Value = 9
)

// Card represents an UNO card.
type Card struct {
	Color Color
	Value Value
}

func (c *Card) String() string {
	return fmt.Sprintf("%d %s", c.Value, c.Color)
}
