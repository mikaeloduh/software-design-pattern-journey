package asciiui

// Text struct represents a text component
type Text struct {
	Position Coordinate
	Text     string
}

// NewText creates a new Text instance
func NewText(position Coordinate, text string) *Text {
	return &Text{
		Position: position,
		Text:     text,
	}
}

// Render delegates the rendering to the theme's RenderText method
func (t *Text) Render(theme Theme) string {
	return theme.RenderText(t)
}

// GetPosition returns the position of the text
func (t *Text) GetPosition() Coordinate {
	return t.Position
}
