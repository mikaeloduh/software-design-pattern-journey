package asciiui

// NumberedList struct represents a numbered list component
type NumberedList struct {
	Position Coordinate
	Lines    []string
}

// NewNumberedList creates a new NumberedList instance
func NewNumberedList(position Coordinate, lines []string) *NumberedList {
	return &NumberedList{
		Position: position,
		Lines:    lines,
	}
}

// Render delegates the rendering to the theme's RenderNumberedList method
func (nl *NumberedList) Render(theme Theme) string {
	return theme.RenderNumberedList(nl)
}

// GetPosition returns the position of the numbered list
func (nl *NumberedList) GetPosition() Coordinate {
	return nl.Position
}
