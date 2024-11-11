package asciiui

import "strings"

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
	return strings.ToUpper(t.renderText(text))
}

// RenderNumberedList renders a numbered list using Roman numerals
func (t *PrettyTheme) RenderNumberedList(list *NumberedList) string {
	return t.renderNumberedList(list, toRoman)
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
