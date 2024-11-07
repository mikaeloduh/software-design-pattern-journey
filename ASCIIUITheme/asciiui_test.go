package asciiui

import (
	"testing"

	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, basicExpected, basicButton.Render(), "Basic Button render failed.")
	assert.Equal(t, prettyExpected, prettyButton.Render(), "Pretty Button render failed.")
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

	assert.Equal(t, basicExpected, basicList.Render(), "Basic NumberedList render failed.")
	assert.Equal(t, prettyExpected, prettyList.Render(), "Pretty NumberedList render failed.")
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

	assert.Equal(t, basicExpected, basicText.Render(), "Basic Text render failed.")
	assert.Equal(t, prettyExpected, prettyText.Render(), "Pretty Text render failed.")
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

	assert.Equal(t, expected, ui.Render(), "UI render failed.")
}

func TestUIRenderPrettyTheme(t *testing.T) {
	ui := NewUI(13, 22) // 調整 UI 的尺寸
	ui.SetTheme(&PrettyThemeFactory{})

	ui.AddComponent(ui.theme.CreateButton(3, 1, "Hi, I miss u", Padding{Width: 1, Height: 0}))
	ui.AddComponent(ui.theme.CreateText(4, 4, "Do u love me ?\nPlease tell..."))
	ui.AddComponent(ui.theme.CreateButton(3, 6, "No", Padding{Width: 1, Height: 0}))
	ui.AddComponent(ui.theme.CreateButton(12, 6, "Yes", Padding{Width: 1, Height: 0}))
	ui.AddComponent(ui.theme.CreateNumberedList(3, 9, []string{"Let's Travel", "Back to home", "Have dinner"}))

	expected := `......................
.  ┌──────────────┐  .
.  │ Hi, I miss u │  .
.  └──────────────┘  .
.   DO U LOVE ME ?   .
.   PLEASE TELL...   .
.  ┌────┐   ┌─────┐  .
.  │ No │   │ Yes │  .
.  └────┘   └─────┘  .
.  i. Let's Travel   .
.  ii. Back to home  .
.  iii. Have dinner  .
......................`

	assert.Equal(t, expected, ui.Render(), "UI render failed for Pretty Theme.")
}
