package asciiui

import "strconv"

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
	return t.renderText(text)
}

// RenderNumberedList renders a numbered list using Arabic numerals
func (t *BasicTheme) RenderNumberedList(list *NumberedList) string {
	return t.renderNumberedList(list, func(n int) string {
		return strconv.Itoa(n)
	})
}
