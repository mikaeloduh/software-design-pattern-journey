package main

import (
	"fmt"

	"asciiuiplugin/asciiui"
)

func main() {
	// Using BasicTheme
	theme := asciiui.NewBasicTheme()
	ui := asciiui.NewUI(22, 13, theme)

	ui.AddComponent(asciiui.NewButton(asciiui.Coordinate{X: 3, Y: 1}, "Hi, I miss u", asciiui.Padding{Width: 1, Height: 0}))
	ui.AddComponent(asciiui.NewText(asciiui.Coordinate{X: 4, Y: 4}, "Do u love me ?\nPlease tell..."))
	ui.AddComponent(asciiui.NewButton(asciiui.Coordinate{X: 3, Y: 6}, "No", asciiui.Padding{Width: 1, Height: 0}))
	ui.AddComponent(asciiui.NewButton(asciiui.Coordinate{X: 12, Y: 6}, "Yes", asciiui.Padding{Width: 1, Height: 0}))
	ui.AddComponent(asciiui.NewNumberedList(asciiui.Coordinate{X: 3, Y: 9}, []string{"Let's Travel", "Back to home", "Have dinner"}))

	fmt.Println(ui.Render())
}
