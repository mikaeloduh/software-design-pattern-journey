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
	Render() []string        // 返回渲染内容的行
	GetPosition() (int, int) // 返回 x 和 y 坐标
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

func (b *BasicButton) Render() []string {
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
	bottom := "+" + strings.Repeat("-", contentWidth) + "+"
	result = append(result, bottom)
	return result
}

func (b *BasicButton) GetPosition() (int, int) {
	return b.x, b.y
}

// PrettyButton 漂亮风格按钮
type PrettyButton struct {
	x, y    int
	text    string
	padding Padding
}

func (p *PrettyButton) Render() []string {
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
	return result
}

func (p *PrettyButton) GetPosition() (int, int) {
	return p.x, p.y
}

// BasicNumberedList 基础风格数字列表
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

// PrettyNumberedList 漂亮风格数字列表
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

// BasicText 基础风格文本
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

// PrettyText 漂亮风格文本
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
	// 初始化画布
	canvas := make([][]rune, ui.height)
	for i := 0; i < ui.height; i++ {
		canvas[i] = make([]rune, ui.width)
		for j := 0; j < ui.width; j++ {
			canvas[i][j] = ' '
		}
	}

	// 绘制边框
	for i := 0; i < ui.height; i++ {
		canvas[i][0] = '.'
		canvas[i][ui.width-1] = '.'
	}
	for j := 0; j < ui.width; j++ {
		canvas[0][j] = '.'
		canvas[ui.height-1][j] = '.'
	}

	// 放置组件
	for _, component := range ui.components {
		x, y := component.GetPosition()
		lines := component.Render()
		for dy, line := range lines {
			row := y + dy
			if row <= 0 || row >= ui.height-1 {
				continue // 超出边界
			}
			for dx, ch := range line {
				col := x + dx
				if col <= 0 || col >= ui.width-1 {
					continue // 超出边界
				}
				canvas[row][col] = ch
			}
		}
	}

	// 生成最终输出
	var output []string
	for _, row := range canvas {
		output = append(output, string(row))
	}
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
