package asciiui

import (
	"fmt"
	"strings"
)

// Theme interface that defines rendering methods for each component type
type Theme interface {
	RenderButton(button *Button) string
	RenderText(text *Text) string
	RenderNumberedList(list *NumberedList) string
}

type BorderStyle struct {
	TopLeft     string
	TopRight    string
	BottomLeft  string
	BottomRight string
	Horizontal  string
	Vertical    string
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

	// Bottom padding
	for i := 0; i < button.Padding.Height; i++ {
		line := border.Vertical + strings.Repeat(" ", totalWidth) + border.Vertical
		middleLines = append(middleLines, line)
	}

	lines := []string{topEdge}
	lines = append(lines, middleLines...)
	lines = append(lines, bottomEdge)

	return strings.Join(lines, "\n")
}

func (bt *BaseTheme) renderText(text *Text) string {
	return text.Text
}

func (bt *BaseTheme) renderNumberedList(list *NumberedList, numberFormatter func(int) string) string {
	var lines []string
	for i, line := range list.Lines {
		number := numberFormatter(i + 1)
		lines = append(lines, fmt.Sprintf("%s. %s", number, line))
	}
	return strings.Join(lines, "\n")
}
