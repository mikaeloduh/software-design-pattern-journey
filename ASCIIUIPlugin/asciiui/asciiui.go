package asciiui

import (
	"fmt"
	"strings"
)

// Component interface that all UI components implement
type Component interface {
	Render(theme Theme) string
	GetPosition() Coordinate // Updated to return Coordinate
}

// Theme interface that defines rendering methods for each component type
type Theme interface {
	RenderButton(button *Button) string
	RenderText(text *Text) string
	RenderNumberedList(list *NumberedList) string
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

// Button struct represents a button component
type Button struct {
	Position Coordinate
	Text     string
	Padding  Padding
}

type BorderStyle struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	Horizontal  string
	Vertical    string
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

type BaseTheme struct {
	Border BorderStyle
}

func (bt *BaseTheme) renderButton(button *Button) string {
	border := bt.Border
	textWidth := len(button.Text)
	totalWidth := 2*button.Padding.Width + textWidth

	topEdge := border.TopLeft + strings.Repeat(border.Horizontal, totalWidth) + border.TopRight
	bottomEdge := border.BottomLeft + strings.Repeat(border.Horizontal, totalWidth) + border.BottomRight

	var middleLines []string

	// Top padding
	for i := 0; i < button.Padding.Height; i++ {
		line := border.Vertical + strings.Repeat(" ", totalWidth) + border.Vertical
		middleLines = append(middleLines, line)
	}

	// Text line
	paddingSpaces := strings.Repeat(" ", button.Padding.Width)
	textLine := border.Vertical + paddingSpaces + button.Text + paddingSpaces + border.Vertical
	middleLines = append(middleLines, textLine)

	// 下邊距
	for i := 0; i < button.Padding.Height; i++ {
		line := border.Vertical + strings.Repeat(" ", totalWidth) + border.Vertical
		middleLines = append(middleLines, line)
	}

	// Bottom padding
	lines := []string{topEdge}
	lines = append(lines, middleLines...)
	lines = append(lines, bottomEdge)

	return strings.Join(lines, "\n")
}

// BasicTheme struct implements the Theme interface with basic ASCII styles
type BasicTheme struct {
	BaseTheme
}

// NewBasicTheme creates a new BasicTheme instance
func NewBasicTheme() *BasicTheme {
	return &BasicTheme{
		BaseTheme{
			Border: BorderStyle{
				TopLeft:     "+",
				TopRight:    "+",
				BottomLeft:  "+",
				BottomRight: "+",
				Horizontal:  "-",
				Vertical:    "|",
			},
		},
	}
}

// RenderButton renders a button using the basic ASCII style
func (t *BasicTheme) RenderButton(button *Button) string {
	return t.renderButton(button)
}

// RenderText renders text without any transformation
func (t *BasicTheme) RenderText(text *Text) string {
	return text.Text
}

// RenderNumberedList renders a numbered list using Arabic numerals
func (t *BasicTheme) RenderNumberedList(list *NumberedList) string {
	var lines []string
	for i, line := range list.Lines {
		lines = append(lines, fmt.Sprintf("%d. %s", i+1, line))
	}
	return strings.Join(lines, "\n")
}

// PrettyTheme struct implements the Theme interface with enhanced ASCII styles
type PrettyTheme struct {
	BaseTheme
}

// NewPrettyTheme creates a new PrettyTheme instance
func NewPrettyTheme() *PrettyTheme {
	return &PrettyTheme{
		BaseTheme{
			Border: BorderStyle{
				TopLeft:     "┌",
				TopRight:    "┐",
				BottomLeft:  "└",
				BottomRight: "┘",
				Horizontal:  "─",
				Vertical:    "│",
			},
		},
	}
}

func (t *PrettyTheme) RenderButton(button *Button) string {
	return t.renderButton(button)
}

// RenderText renders text in uppercase
func (t *PrettyTheme) RenderText(text *Text) string {
	return strings.ToUpper(text.Text)
}

// RenderNumberedList renders a numbered list using Roman numerals
func (t *PrettyTheme) RenderNumberedList(list *NumberedList) string {
	var lines []string
	for i, line := range list.Lines {
		lines = append(lines, fmt.Sprintf("%s. %s", toRoman(i+1), line))
	}
	return strings.Join(lines, "\n")
}

// Helper function to convert integer to Roman numerals
func toRoman(num int) string {
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syb := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var roman strings.Builder
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman.WriteString(syb[i])
		}
	}
	return roman.String()
}

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
