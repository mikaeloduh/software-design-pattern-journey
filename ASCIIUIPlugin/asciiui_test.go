package asciiui

import "testing"

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
