package asciiui

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Padding defines the padding for components
type Padding struct {
	Width  int
	Height int
}

// AsciiComponent interface defines the common behavior of all components
type AsciiComponent interface {
	Render() []string        // Returns the rendered lines of the component
	GetPosition() (int, int) // Returns the x and y coordinates
}

// AsciiUIFactory is the abstract factory interface
type AsciiUIFactory interface {
	CreateButton(x, y int, text string, padding Padding) AsciiComponent
	CreateNumberedList(x, y int, lines []string) AsciiComponent
	CreateText(x, y int, text string) AsciiComponent
}

// BasicThemeFactory is the factory for the basic ASCII theme
type BasicThemeFactory struct{}

func (b *BasicThemeFactory) CreateButton(x, y int, text string, padding Padding) AsciiComponent {
	return &BasicButton{x, y, text, padding}
}

func (b *BasicThemeFactory) CreateNumberedList(x, y int, lines []string) AsciiComponent {
	return &BasicNumberedList{x, y, lines}
}

func (b *BasicThemeFactory) CreateText(x, y int, text string) AsciiComponent {
	return &BasicText{x, y, text}
}

// PrettyThemeFactory is the factory for the pretty ASCII theme
type PrettyThemeFactory struct{}

func (p *PrettyThemeFactory) CreateButton(x, y int, text string, padding Padding) AsciiComponent {
	return &PrettyButton{x, y, text, padding}
}

func (p *PrettyThemeFactory) CreateNumberedList(x, y int, lines []string) AsciiComponent {
	return &PrettyNumberedList{x, y, lines}
}

func (p *PrettyThemeFactory) CreateText(x, y int, text string) AsciiComponent {
	return &PrettyText{x, y, text}
}

// BasicButton represents a button in the basic theme
type BasicButton struct {
	x, y    int
	text    string
	padding Padding
}

func (b *BasicButton) Render() []string {
	pw, ph := b.padding.Width, b.padding.Height
	textLength := utf8.RuneCountInString(b.text)
	contentWidth := textLength + pw*2
	top := "+" + strings.Repeat("-", contentWidth) + "+"
	side := "|" + strings.Repeat(" ", contentWidth) + "|"
	content := "|" + strings.Repeat(" ", pw) + b.text + strings.Repeat(" ", pw) + "|"

	var result []string
	result = append(result, top)
	for i := 0; i < ph; i++ {
		result = append(result, side)
	}
	result = append(result, content)
	for i := 0; i < ph; i++ {
		result = append(result, side)
	}
	bottom := "+" + strings.Repeat("-", contentWidth) + "+"
	result = append(result, bottom)
	return result
}

func (b *BasicButton) GetPosition() (int, int) {
	return b.x, b.y
}

// PrettyButton represents a button in the pretty theme
type PrettyButton struct {
	x, y    int
	text    string
	padding Padding
}

func (p *PrettyButton) Render() []string {
	pw, ph := p.padding.Width, p.padding.Height
	textLength := utf8.RuneCountInString(p.text)
	contentWidth := textLength + pw*2
	top := "┌" + strings.Repeat("─", contentWidth) + "┐"
	side := "│" + strings.Repeat(" ", contentWidth) + "│"
	content := "│" + strings.Repeat(" ", pw) + p.text + strings.Repeat(" ", pw) + "│"

	var result []string
	result = append(result, top)
	for i := 0; i < ph; i++ {
		result = append(result, side)
	}
	result = append(result, content)
	for i := 0; i < ph; i++ {
		result = append(result, side)
	}
	bottom := "└" + strings.Repeat("─", contentWidth) + "┘"
	result = append(result, bottom)
	return result
}

func (p *PrettyButton) GetPosition() (int, int) {
	return p.x, p.y
}

// BasicNumberedList represents a numbered list in the basic theme
type BasicNumberedList struct {
	x, y  int
	lines []string
}

func (b *BasicNumberedList) Render() []string {
	var result []string
	for i, line := range b.lines {
		result = append(result, fmt.Sprintf("%d. %s", i+1, line))
	}
	return result
}

func (b *BasicNumberedList) GetPosition() (int, int) {
	return b.x, b.y
}

// PrettyNumberedList represents a numbered list in the pretty theme
type PrettyNumberedList struct {
	x, y  int
	lines []string
}

func (p *PrettyNumberedList) Render() []string {
	var result []string
	for i, line := range p.lines {
		roman := toRoman(i + 1)
		result = append(result, fmt.Sprintf("%s. %s", roman, line))
	}
	return result
}

func (p *PrettyNumberedList) GetPosition() (int, int) {
	return p.x, p.y
}

// BasicText represents text in the basic theme
type BasicText struct {
	x, y int
	text string
}

func (b *BasicText) Render() []string {
	return strings.Split(b.text, "\n")
}

func (b *BasicText) GetPosition() (int, int) {
	return b.x, b.y
}

// PrettyText represents text in the pretty theme
type PrettyText struct {
	x, y int
	text string
}

func (p *PrettyText) Render() []string {
	upperText := strings.ToUpper(p.text)
	return strings.Split(upperText, "\n")
}

func (p *PrettyText) GetPosition() (int, int) {
	return p.x, p.y
}

// UI represents the ASCII UI
type UI struct {
	height, width int
	theme         AsciiUIFactory
	components    []AsciiComponent
}

func NewUI(height, width int) *UI {
	return &UI{
		height:     height,
		width:      width,
		components: []AsciiComponent{},
	}
}

func (ui *UI) SetTheme(theme AsciiUIFactory) {
	ui.theme = theme
}

func (ui *UI) AddComponent(component AsciiComponent) {
	ui.components = append(ui.components, component)
}

func (ui *UI) Render() string {
	// Initialize the canvas
	canvas := make([][]rune, ui.height)
	for i := 0; i < ui.height; i++ {
		canvas[i] = make([]rune, ui.width)
		for j := 0; j < ui.width; j++ {
			canvas[i][j] = ' '
		}
	}

	// Draw the borders
	for i := 0; i < ui.height; i++ {
		canvas[i][0] = '.'
		canvas[i][ui.width-1] = '.'
	}
	for j := 0; j < ui.width; j++ {
		canvas[0][j] = '.'
		canvas[ui.height-1][j] = '.'
	}

	// Place the components
	for _, component := range ui.components {
		x, y := component.GetPosition()
		lines := component.Render()
		for dy, line := range lines {
			row := y + dy
			if row <= 0 || row >= ui.height-1 {
				continue // Out of bounds
			}
			col := x
			for _, ch := range line {
				if col <= 0 || col >= ui.width-1 {
					break // Out of bounds
				}
				canvas[row][col] = ch
				col += 1
			}
		}
	}

	// Generate the final output
	var output []string
	for _, row := range canvas {
		output = append(output, string(row))
	}
	return strings.Join(output, "\n")
}

// Helper function to convert integers to lowercase Roman numerals
func toRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var result strings.Builder
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			num -= vals[i]
			result.WriteString(syms[i])
		}
	}
	return strings.ToLower(result.String())
}
