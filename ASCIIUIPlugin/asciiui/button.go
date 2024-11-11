package asciiui

// Button struct represents a button component
type Button struct {
	Position Coordinate
	Text     string
	Padding  Padding
}

// NewButton creates a new Button instance
func NewButton(position Coordinate, text string, padding Padding) *Button {
	return &Button{
		Position: position,
		Text:     text,
		Padding:  padding,
	}
}

// Render delegates the rendering to the theme's RenderButton method
func (b *Button) Render(theme Theme) string {
	return theme.RenderButton(b)
}

// GetPosition returns the position of the button
func (b *Button) GetPosition() Coordinate {
	return b.Position
}
