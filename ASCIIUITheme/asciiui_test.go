package asciiui

import (
	"testing"
)

func TestButtonRender(t *testing.T) {
	basicFactory := &BasicThemeFactory{}
	prettyFactory := &PrettyThemeFactory{}

	basicButton := basicFactory.CreateButton(0, 0, "Hi, I miss u", Padding{Width: 1, Height: 0})
	prettyButton := prettyFactory.CreateButton(0, 0, "Hi, I miss u", Padding{Width: 1, Height: 0})

	basicExpected := "+--------------+\n| Hi, I miss u |\n+--------------+"
	prettyExpected := "┌──────────────┐\n│ Hi, I miss u │\n└──────────────┘"

	if basicButton.Render() != basicExpected {
		t.Errorf("Basic Button render failed.\nExpected:\n%s\nGot:\n%s", basicExpected, basicButton.Render())
	}

	if prettyButton.Render() != prettyExpected {
		t.Errorf("Pretty Button render failed.\nExpected:\n%s\nGot:\n%s", prettyExpected, prettyButton.Render())
	}
}

func TestNumberedListRender(t *testing.T) {
	basicFactory := &BasicThemeFactory{}
	prettyFactory := &PrettyThemeFactory{}

	lines := []string{"Let's Travel", "Back to home", "Have dinner"}

	basicList := basicFactory.CreateNumberedList(0, 0, lines)
	prettyList := prettyFactory.CreateNumberedList(0, 0, lines)

	basicExpected := "1. Let's Travel\n2. Back to home\n3. Have dinner"
	prettyExpected := "i. Let's Travel\nii. Back to home\niii. Have dinner"

	if basicList.Render() != basicExpected {
		t.Errorf("Basic NumberedList render failed.\nExpected:\n%s\nGot:\n%s", basicExpected, basicList.Render())
	}

	if prettyList.Render() != prettyExpected {
		t.Errorf("Pretty NumberedList render failed.\nExpected:\n%s\nGot:\n%s", prettyExpected, prettyList.Render())
	}
}

func TestTextRender(t *testing.T) {
	basicFactory := &BasicThemeFactory{}
	prettyFactory := &PrettyThemeFactory{}

	textContent := "Do u love me ?\nPlease tell..."

	basicText := basicFactory.CreateText(0, 0, textContent)
	prettyText := prettyFactory.CreateText(0, 0, textContent)

	basicExpected := "Do u love me ?\nPlease tell..."
	prettyExpected := "DO U LOVE ME ?\nPLEASE TELL..."

	if basicText.Render() != basicExpected {
		t.Errorf("Basic Text render failed.\nExpected:\n%s\nGot:\n%s", basicExpected, basicText.Render())
	}

	if prettyText.Render() != prettyExpected {
		t.Errorf("Pretty Text render failed.\nExpected:\n%s\nGot:\n%s", prettyExpected, prettyText.Render())
	}
}

func TestUIRender(t *testing.T) {
	ui := NewUI(22, 22)
	ui.SetTheme(&BasicThemeFactory{})

	ui.AddComponent(ui.theme.CreateButton(3, 1, "Hi, I miss u", Padding{Width: 1, Height: 0}))
	ui.AddComponent(ui.theme.CreateText(4, 4, "Do u love me ?\nPlease tell..."))
	ui.AddComponent(ui.theme.CreateButton(3, 6, "No", Padding{Width: 1, Height: 0}))
	ui.AddComponent(ui.theme.CreateButton(12, 6, "Yes", Padding{Width: 1, Height: 0}))
	ui.AddComponent(ui.theme.CreateNumberedList(3, 9, []string{"Let's Travel", "Back to home", "Have dinner"}))

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
......................`

	if ui.Render() != expected {
		t.Errorf("UI render failed.\nExpected:\n%s\nGot:\n%s", expected, ui.Render())
	}
}
