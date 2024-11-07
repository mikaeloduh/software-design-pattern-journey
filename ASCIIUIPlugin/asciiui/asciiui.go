package asciiui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mattn/go-runewidth"
)

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

	//textRunes := []rune(b.Text)
	textWidth := runewidth.StringWidth(b.Text)
	totalWidth := 2*b.Padding.Width + textWidth
	//totalHeight := 2*b.Padding.Height + 1

	// Repeat horizontal edge without spaces
	topEdge := tl + strings.Repeat(hEdge, totalWidth) + tr
	bottomEdge := bl + strings.Repeat(hEdge, totalWidth) + br

	var middleLines []string

	// Top padding
	for i := 0; i < b.Padding.Height; i++ {
		line := vEdge + strings.Repeat(" ", totalWidth) + vEdge
		middleLines = append(middleLines, line)
	}

	// Text line
	paddingSpaces := strings.Repeat(" ", b.Padding.Width)
	textLine := vEdge + paddingSpaces + b.Text + paddingSpaces + vEdge
	middleLines = append(middleLines, textLine)

	// Bottom padding
	for i := 0; i < b.Padding.Height; i++ {
		line := vEdge + strings.Repeat(" ", totalWidth) + vEdge
		middleLines = append(middleLines, line)
	}

	lines := []string{topEdge}
	lines = append(lines, middleLines...)
	lines = append(lines, bottomEdge)

	return strings.Join(lines, "\n")
}

type NumberedList struct {
	X     int
	Y     int
	Lines []string
}

func NewNumberedList(x, y int, lines []string) *NumberedList {
	return &NumberedList{
		X:     x,
		Y:     y,
		Lines: lines,
	}
}

func (nl *NumberedList) Render(theme Theme) string {
	var result []string
	for i, line := range nl.Lines {
		prefixTemplate := theme.GetStyle("number.prefix")
		var formattedNumber string
		if strings.Contains(prefixTemplate, "%s") {
			formattedNumber = fmt.Sprintf(prefixTemplate, romanNumeral(i+1))
		} else {
			formattedNumber = fmt.Sprintf(prefixTemplate, i+1)
		}
		result = append(result, formattedNumber+line)
	}
	return strings.Join(result, "\n")
}

func romanNumeral(num int) string {
	numerals := []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	if num < 10 {
		return numerals[num]
	}
	// 简单处理，只支持到 9
	return strconv.Itoa(num)
}

type Text struct {
	X    int
	Y    int
	Text string
}

func NewText(x, y int, text string) *Text {
	return &Text{
		X:    x,
		Y:    y,
		Text: text,
	}
}

func (t *Text) Render(theme Theme) string {
	transform := theme.GetStyle("text.transform")
	lines := strings.Split(t.Text, "\n")
	for i, line := range lines {
		if transform == "upper" {
			lines[i] = strings.ToUpper(line)
		} else {
			lines[i] = line
		}
	}
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
			"number.prefix":          "%d. ",
			"text.transform":         "none",
		},
	}
}

func (t *BasicTheme) GetStyle(styleKey string) string {
	if style, ok := t.Styles[styleKey]; ok {
		return style
	}
	return ""
}

type PrettyTheme struct {
	Styles map[string]string
}

func NewPrettyTheme() *PrettyTheme {
	return &PrettyTheme{
		Styles: map[string]string{
			"button.corner.tl":       "┌",
			"button.corner.tr":       "┐",
			"button.corner.bl":       "└",
			"button.corner.br":       "┘",
			"button.edge.horizontal": "─",
			"button.edge.vertical":   "│",
			"number.prefix":          "%s. ",
			"text.transform":         "upper",
		},
	}
}

func (t *PrettyTheme) GetStyle(styleKey string) string {
	if style, ok := t.Styles[styleKey]; ok {
		return style
	}
	return ""
}

type UI struct {
	Width      int
	Height     int
	Components []Component
	Theme      Theme
}

func NewUI(width, height int, theme Theme) *UI {
	return &UI{
		Width:      width,
		Height:     height,
		Theme:      theme,
		Components: []Component{},
	}
}

func (ui *UI) AddComponent(c Component) {
	ui.Components = append(ui.Components, c)
}

func (ui *UI) Render() string {
	// 初始化画布，使用空格填充
	canvas := make([][]rune, ui.Height)
	for i := range canvas {
		canvas[i] = make([]rune, ui.Width)
		for j := range canvas[i] {
			canvas[i][j] = ' '
		}
	}

	// 绘制 UI 边框
	for i := 0; i < ui.Width; i++ {
		canvas[0][i] = '.'
		canvas[ui.Height-1][i] = '.'
	}
	for i := 0; i < ui.Height; i++ {
		canvas[i][0] = '.'
		canvas[i][ui.Width-1] = '.'
	}

	// 渲染每个组件并放置到画布上
	for _, c := range ui.Components {
		rendered := c.Render(ui.Theme)
		lines := strings.Split(rendered, "\n")
		// 获取组件的位置
		var x, y int
		switch comp := c.(type) {
		case *Button:
			x, y = comp.X, comp.Y
		case *Text:
			x, y = comp.X, comp.Y
		case *NumberedList:
			x, y = comp.X, comp.Y
		}

		// 将组件的渲染结果放置到画布上
		for i, line := range lines {
			canvasY := y + i
			if canvasY <= 0 || canvasY >= ui.Height-1 {
				continue
			}
			lineRunes := []rune(line)
			for j, ch := range lineRunes {
				canvasX := x + runewidth.StringWidth(string(lineRunes[:j]))
				if canvasX <= 0 || canvasX >= ui.Width-1 {
					continue
				}
				canvas[canvasY][canvasX] = ch
			}
		}
	}

	// 将画布转换为字符串
	var builder strings.Builder
	for _, line := range canvas {
		builder.WriteString(string(line))
		builder.WriteString("\n")
	}

	return builder.String()
}
