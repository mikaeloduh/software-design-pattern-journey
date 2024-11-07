package asciiui

import (
	"testing"
)

func TestButtonRender(t *testing.T) {
	basicFactory := &BasicThemeFactory{}
	prettyFactory := &PrettyThemeFactory{}

	basicButton := basicFactory.CreateButton(0, 0, "Hi, I miss u", Padding{Width: 1, Height: 0})
	prettyButton := prettyFactory.CreateButton(0, 0, "Hi, I miss u", Padding{Width: 1, Height: 0})

	basicExpected := []string{
		"+--------------+",
		"| Hi, I miss u |",
		"+--------------+",
	}
	prettyExpected := []string{
		"┌──────────────┐",
		"│ Hi, I miss u │",
		"└──────────────┘",
	}

	if !compareSlices(basicButton.Render(), basicExpected) {
		t.Errorf("Basic Button render failed.\nExpected:\n%v\nGot:\n%v", basicExpected, basicButton.Render())
	}

	if !compareSlices(prettyButton.Render(), prettyExpected) {
		t.Errorf("Pretty Button render failed.\nExpected:\n%v\nGot:\n%v", prettyExpected, prettyButton.Render())
	}
}

func TestNumberedListRender(t *testing.T) {
	basicFactory := &BasicThemeFactory{}
	prettyFactory := &PrettyThemeFactory{}

	lines := []string{"Let's Travel", "Back to home", "Have dinner"}

	basicList := basicFactory.CreateNumberedList(0, 0, lines)
	prettyList := prettyFactory.CreateNumberedList(0, 0, lines)

	basicExpected := []string{
		"1. Let's Travel",
		"2. Back to home",
		"3. Have dinner",
	}
	prettyExpected := []string{
		"i. Let's Travel",
		"ii. Back to home",
		"iii. Have dinner",
	}

	if !compareSlices(basicList.Render(), basicExpected) {
		t.Errorf("Basic NumberedList render failed.\nExpected:\n%v\nGot:\n%v", basicExpected, basicList.Render())
	}

	if !compareSlices(prettyList.Render(), prettyExpected) {
		t.Errorf("Pretty NumberedList render failed.\nExpected:\n%v\nGot:\n%v", prettyExpected, prettyList.Render())
	}
}

func TestTextRender(t *testing.T) {
	basicFactory := &BasicThemeFactory{}
	prettyFactory := &PrettyThemeFactory{}

	textContent := "Do u love me ?\nPlease tell..."

	basicText := basicFactory.CreateText(0, 0, textContent)
	prettyText := prettyFactory.CreateText(0, 0, textContent)

	basicExpected := []string{
		"Do u love me ?",
		"Please tell...",
	}
	prettyExpected := []string{
		"DO U LOVE ME ?",
		"PLEASE TELL...",
	}

	if !compareSlices(basicText.Render(), basicExpected) {
		t.Errorf("Basic Text render failed.\nExpected:\n%v\nGot:\n%v", basicExpected, basicText.Render())
	}

	if !compareSlices(prettyText.Render(), prettyExpected) {
		t.Errorf("Pretty Text render failed.\nExpected:\n%v\nGot:\n%v", prettyExpected, prettyText.Render())
	}
}

func TestUIRender(t *testing.T) {
	ui := NewUI(13, 22)
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

func compareSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
