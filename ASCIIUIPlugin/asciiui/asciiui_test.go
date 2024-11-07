package asciiui

import (
	"testing"
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
	if rendered != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, rendered)
	}
}

func TestNumberedListRenderingWithBasicTheme(t *testing.T) {
	theme := NewBasicTheme()
	list := NewNumberedList(0, 0, []string{"Apple", "Banana", "Grape"})
	rendered := list.Render(theme)
	expected := `1. Apple
2. Banana
3. Grape`
	if rendered != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, rendered)
	}
}

func TestTextRenderingWithPrettyTheme(t *testing.T) {
	theme := NewPrettyTheme()
	text := NewText(0, 0, "Do u love me ?\nPlease tell...")
	rendered := text.Render(theme)
	expected := "DO U LOVE ME ?\nPLEASE TELL..."
	if rendered != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, rendered)
	}
}

func TestUIRenderingWithBasicTheme(t *testing.T) {
	theme := NewBasicTheme()
	ui := NewUI(22, 13, theme) // 修改宽度为22，高度为13

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
	if rendered != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, rendered)
	}
}

func TestUIRenderingWithPrettyTheme(t *testing.T) {
	theme := NewPrettyTheme()
	ui := NewUI(22, 13, theme) // 修改宽度为22，高度为13

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
	if rendered != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, rendered)
	}
}
