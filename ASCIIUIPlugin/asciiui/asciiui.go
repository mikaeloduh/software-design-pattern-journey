package asciiui

import (
	"strings"
)

// UI struct represents the user interface containing components and a theme
type UI struct {
	Width      int
	Height     int
	Components []Component
	Theme      Theme
}

// NewUI creates a new UI instance with the specified dimensions and theme
func NewUI(width, height int, theme Theme) *UI {
	return &UI{
		Width:      width,
		Height:     height,
		Components: []Component{},
		Theme:      theme,
	}
}

// AddComponent adds a component to the UI
func (ui *UI) AddComponent(c Component) {
	ui.Components = append(ui.Components, c)
}

// SetTheme allows changing the theme of the UI
func (ui *UI) SetTheme(theme Theme) {
	ui.Theme = theme
}

// Render renders the UI by placing components onto a canvas
func (ui *UI) Render() string {
	// Initialize canvas filled with spaces
	canvas := make([][]rune, ui.Height)
	for i := range canvas {
		canvas[i] = make([]rune, ui.Width)
		for j := range canvas[i] {
			canvas[i][j] = ' '
		}
	}

	// Draw UI borders
	for i := 0; i < ui.Width; i++ {
		canvas[0][i] = '.'
		canvas[ui.Height-1][i] = '.'
	}
	for i := 0; i < ui.Height; i++ {
		canvas[i][0] = '.'
		canvas[i][ui.Width-1] = '.'
	}

	// Render each component and place it onto the canvas
	for _, c := range ui.Components {
		rendered := c.Render(ui.Theme)
		lines := strings.Split(rendered, "\n")
		// Get component position using GetPosition()
		pos := c.GetPosition()
		x, y := pos.X, pos.Y
		for i, line := range lines {
			canvasY := y + i
			if canvasY <= 0 || canvasY >= ui.Height-1 {
				continue
			}
			lineRunes := []rune(line)
			for j, ch := range lineRunes {
				canvasX := x + j
				if canvasX <= 0 || canvasX >= ui.Width-1 {
					continue
				}
				canvas[canvasY][canvasX] = ch
			}
		}
	}

	// Convert canvas to string
	var builder strings.Builder
	for _, line := range canvas {
		builder.WriteString(string(line))
		builder.WriteString("\n")
	}

	return builder.String()
}
