package asciiui

// Component interface that all UI components implement
type Component interface {
	Render(theme Theme) string
	GetPosition() Coordinate // Updated to return Coordinate
}

// Coordinate struct represents the position of a component
type Coordinate struct {
	X int
	Y int
}

// Padding struct defines the padding around the button text
type Padding struct {
	Width  int
	Height int
}
