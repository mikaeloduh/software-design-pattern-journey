package asciiui

import "strings"

type Component interface {
	Render(theme Theme) string
}

type Theme interface {
	GetStyle(styleKey string) string
}

type Padding struct {
	Width  int
	Height int
}

type Button struct {
	X       int
	Y       int
	Text    string
	Padding Padding
}

func NewButton(x, y int, text string, padding Padding) *Button {
	return &Button{
		X:       x,
		Y:       y,
		Text:    text,
		Padding: padding,
	}
}

func (b *Button) Render(theme Theme) string {
	tl := theme.GetStyle("button.corner.tl")
	tr := theme.GetStyle("button.corner.tr")
	bl := theme.GetStyle("button.corner.bl")
	br := theme.GetStyle("button.corner.br")
	hEdge := theme.GetStyle("button.edge.horizontal")
	vEdge := theme.GetStyle("button.edge.vertical")

	textWidth := len(b.Text)
	totalWidth := 2*b.Padding.Width + textWidth
	//totalHeight := 2*b.Padding.Height + 1

	topEdge := tl + strings.Repeat(hEdge, totalWidth) + tr
	bottomEdge := bl + strings.Repeat(hEdge, totalWidth) + br

	var middleLines []string

	for i := 0; i < b.Padding.Height; i++ {
		line := vEdge + strings.Repeat(" ", totalWidth) + vEdge
		middleLines = append(middleLines, line)
	}

	textLine := vEdge + strings.Repeat(" ", b.Padding.Width) + b.Text + strings.Repeat(" ", b.Padding.Width) + vEdge
	middleLines = append(middleLines, textLine)

	for i := 0; i < b.Padding.Height; i++ {
		line := vEdge + strings.Repeat(" ", totalWidth) + vEdge
		middleLines = append(middleLines, line)
	}

	lines := []string{topEdge}
	lines = append(lines, middleLines...)
	lines = append(lines, bottomEdge)

	return strings.Join(lines, "\n")
}

type BasicTheme struct {
	Styles map[string]string
}

func NewBasicTheme() *BasicTheme {
	return &BasicTheme{
		Styles: map[string]string{
			"button.corner.tl":       "+",
			"button.corner.tr":       "+",
			"button.corner.bl":       "+",
			"button.corner.br":       "+",
			"button.edge.horizontal": "-",
			"button.edge.vertical":   "|",
		},
	}
}

func (t *BasicTheme) GetStyle(styleKey string) string {
	if style, ok := t.Styles[styleKey]; ok {
		return style
	}
	return ""
}
