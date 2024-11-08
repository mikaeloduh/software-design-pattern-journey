package asciiui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestButtonRenderingWithBasicTheme(t *testing.T) {
	theme := NewBasicTheme()
	button := NewButton(0, 0, "Example", Padding{Width: 3, Height: 1})
	rendered := button.Render(theme)
	expected := `+-------------+
|             |
|   Example   |
|             |
+-------------+`
	assert.Equal(t, expected, rendered, "Button rendering with basic theme should match expected output")
}

func TestNumberedListRenderingWithBasicTheme(t *testing.T) {
	theme := NewBasicTheme()
	list := NewNumberedList(0, 0, []string{"Apple", "Banana", "Grape"})
	rendered := list.Render(theme)
	expected := `1. Apple
2. Banana
3. Grape`
	assert.Equal(t, expected, rendered, "Numbered list rendering with basic theme should match expected output")
}

func TestTextRenderingWithPrettyTheme(t *testing.T) {
	theme := NewPrettyTheme()
	text := NewText(0, 0, "Do u love me ?\nPlease tell...")
	rendered := text.Render(theme)
	expected := "DO U LOVE ME ?\nPLEASE TELL..."
	assert.Equal(t, expected, rendered, "Text rendering with pretty theme should match expected output")
}

func TestUIRenderingWithBasicTheme(t *testing.T) {
	theme := NewBasicTheme()
	ui := NewUI(22, 13, theme)

	ui.AddComponent(NewButton(3, 1, "Hi, I miss u", Padding{Width: 1, Height: 0}))
	ui.AddComponent(NewText(4, 4, "Do u love me ?\nPlease tell..."))
	ui.AddComponent(NewButton(3, 6, "No", Padding{Width: 1, Height: 0}))
	ui.AddComponent(NewButton(12, 6, "Yes", Padding{Width: 1, Height: 0}))
	ui.AddComponent(NewNumberedList(3, 9, []string{"Let's Travel", "Back to home", "Have dinner"}))

	rendered := ui.Render()
	expected := `......................
.  +--------------+  .
.  | Hi, I miss u |  .
.  +--------------+  .
.   Do u love me ?   .
.   Please tell...   .
.  +----+   +-----+  .
.  | No |   | Yes |  .
.  +----+   +-----+  .
.  1. Let's Travel   .
.  2. Back to home   .
.  3. Have dinner    .
......................
`
	assert.Equal(t, expected, rendered, "UI rendering with basic theme should match expected output")
}

func TestUIRenderingWithPrettyTheme(t *testing.T) {
	theme := NewPrettyTheme()
	ui := NewUI(22, 13, theme)

	ui.AddComponent(NewButton(3, 1, "Hi, I miss u", Padding{Width: 1, Height: 0}))
	ui.AddComponent(NewText(4, 4, "Do u love me ?\nPlease tell..."))
	ui.AddComponent(NewButton(3, 6, "No", Padding{Width: 1, Height: 0}))
	ui.AddComponent(NewButton(12, 6, "Yes", Padding{Width: 1, Height: 0}))
	ui.AddComponent(NewNumberedList(3, 9, []string{"Let's Travel", "Back to home", "Have dinner"}))

	rendered := ui.Render()
	expected := `......................
.  ┌──────────────┐  .
.  │ Hi, I miss u │  .
.  └──────────────┘  .
.   DO U LOVE ME ?   .
.   PLEASE TELL...   .
.  ┌────┐   ┌─────┐  .
.  │ No │   │ Yes │  .
.  └────┘   └─────┘  .
.  I. Let's Travel   .
.  II. Back to home  .
.  III. Have dinner  .
......................
`
	assert.Equal(t, expected, rendered, "UI rendering with pretty theme should match expected output")
}
