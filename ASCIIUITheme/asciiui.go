package asciiui

import (
	"fmt"
	"strings"
)

// Padding 定义内距
type Padding struct {
	Width  int
	Height int
}

// AsciiComponent 定义所有控件的通用接口
type AsciiComponent interface {
	Render() string
}

// AsciiUIFactory 抽象工厂接口
type AsciiUIFactory interface {
	CreateButton(x, y int, text string, padding Padding) AsciiComponent
	CreateNumberedList(x, y int, lines []string) AsciiComponent
	CreateText(x, y int, text string) AsciiComponent
}

// BasicThemeFactory 基础风格工厂
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

// PrettyThemeFactory 漂亮风格工厂
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

// BasicButton 基础风格按钮
type BasicButton struct {
	x, y    int
	text    string
	padding Padding
}

func (b *BasicButton) Render() string {
	pw, ph := b.padding.Width, b.padding.Height
	contentWidth := len(b.text) + pw*2
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
	result = append(result, top)
	return strings.Join(result, "\n")
}

// PrettyButton 漂亮风格按钮
type PrettyButton struct {
	x, y    int
	text    string
	padding Padding
}

func (p *PrettyButton) Render() string {
	pw, ph := p.padding.Width, p.padding.Height
	contentWidth := len(p.text) + pw*2
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
	return strings.Join(result, "\n")
}

// BasicNumberedList 基础风格数字列表
type BasicNumberedList struct {
	x, y  int
	lines []string
}

func (b *BasicNumberedList) Render() string {
	var result []string
	for i, line := range b.lines {
		result = append(result, fmt.Sprintf("%d. %s", i+1, line))
	}
	return strings.Join(result, "\n")
}

// PrettyNumberedList 漂亮风格数字列表
type PrettyNumberedList struct {
	x, y  int
	lines []string
}

func (p *PrettyNumberedList) Render() string {
	var result []string
	for i, line := range p.lines {
		roman := toRoman(i + 1)
		result = append(result, fmt.Sprintf("%s. %s", roman, line))
	}
	return strings.Join(result, "\n")
}

// BasicText 基础风格文本
type BasicText struct {
	x, y int
	text string
}

func (b *BasicText) Render() string {
	return b.text
}

// PrettyText 漂亮风格文本
type PrettyText struct {
	x, y int
	text string
}

func (p *PrettyText) Render() string {
	return strings.ToUpper(p.text)
}

// UI 界面结构
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
	var output []string
	border := strings.Repeat(".", ui.width)
	output = append(output, border)

	for _, component := range ui.components {
		lines := strings.Split(component.Render(), "\n")
		for _, line := range lines {
			formattedLine := fmt.Sprintf(".  %-18s.", line)
			output = append(output, formattedLine)
		}
	}

	output = append(output, border)
	return strings.Join(output, "\n")
}

// 辅助函数：整数转罗马数字（小写）
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
