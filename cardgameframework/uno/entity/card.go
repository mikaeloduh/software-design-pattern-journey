package entity

import (
	"fmt"
)

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

//type IUnoCard interface {
//	template.ICard
//}

// UnoCard represents an UNO card.
type UnoCard struct {
	Color Color
	Value Value
}

func (c UnoCard) String() string {
	return fmt.Sprintf("%d %s", c.Value, c.Color)
}

func (c UnoCard) Compare(other UnoCard) int {
	if c.Color == other.Color || c.Value == other.Value {
		return 0
	}
	return -1
}
