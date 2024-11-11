package asciiui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestButtonRendering tests the Button component rendering with both themes
func TestButtonRendering(t *testing.T) {
	button := NewButton(Coordinate{X: 0, Y: 0}, "Click Me", Padding{Width: 1, Height: 0})

	t.Run("BasicTheme", func(t *testing.T) {
		theme := NewBasicTheme()
		rendered := button.Render(theme)
		expected := `+----------+
| Click Me |
+----------+`

		assert.Equal(t, expected, rendered, "Button rendering with BasicTheme should match expected output")
	})

	t.Run("PrettyTheme", func(t *testing.T) {
		theme := NewPrettyTheme()
		rendered := button.Render(theme)
		expected := `┌──────────┐
│ Click Me │
└──────────┘`

		assert.Equal(t, expected, rendered, "Button rendering with PrettyTheme should match expected output")
	})
}

// TestTextRendering tests the Text component rendering with both themes
func TestTextRendering(t *testing.T) {
	text := NewText(Coordinate{X: 0, Y: 0}, "Hello, World!")

	t.Run("BasicTheme", func(t *testing.T) {
		theme := NewBasicTheme()
		rendered := text.Render(theme)
		expected := "Hello, World!"

		assert.Equal(t, expected, rendered, "Text rendering with BasicTheme should match expected output")
	})

	t.Run("PrettyTheme", func(t *testing.T) {
		theme := NewPrettyTheme()
		rendered := text.Render(theme)
		expected := "HELLO, WORLD!"

		assert.Equal(t, expected, rendered, "Text rendering with PrettyTheme should match expected output")
	})
}

// TestNumberedListRendering tests the NumberedList component rendering with both themes
func TestNumberedListRendering(t *testing.T) {
	list := NewNumberedList(Coordinate{X: 0, Y: 0}, []string{"Apple", "Banana", "Cherry"})

	t.Run("BasicTheme", func(t *testing.T) {
		theme := NewBasicTheme()
		rendered := list.Render(theme)
		expected := `1. Apple
2. Banana
3. Cherry`

		assert.Equal(t, expected, rendered, "NumberedList rendering with BasicTheme should match expected output")
	})

	t.Run("PrettyTheme", func(t *testing.T) {
		theme := NewPrettyTheme()
		rendered := list.Render(theme)
		expected := `I. Apple
II. Banana
III. Cherry`

		assert.Equal(t, expected, rendered, "NumberedList rendering with PrettyTheme should match expected output")
	})
}

func TestUIRendering(t *testing.T) {
	ui := NewUI(22, 13, NewBasicTheme())

	ui.AddComponent(NewButton(Coordinate{X: 3, Y: 1}, "Hi, I miss u", Padding{Width: 1, Height: 0}))
	ui.AddComponent(NewText(Coordinate{X: 4, Y: 4}, "Do u love me ?\nPlease tell..."))
	ui.AddComponent(NewButton(Coordinate{X: 3, Y: 6}, "No", Padding{Width: 1, Height: 0}))
	ui.AddComponent(NewButton(Coordinate{X: 12, Y: 6}, "Yes", Padding{Width: 1, Height: 0}))
	ui.AddComponent(NewNumberedList(Coordinate{X: 3, Y: 9}, []string{"Let's Travel", "Back to home", "Have dinner"}))

	t.Run("BasicTheme", func(t *testing.T) {
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
	})

	t.Run("PrettyTheme", func(t *testing.T) {
		// Change to PrettyTheme using SetTheme
		ui.SetTheme(NewPrettyTheme())

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
	})
}
