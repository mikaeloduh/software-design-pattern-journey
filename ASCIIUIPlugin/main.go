package main

import "fmt"
import "asciiuiplugin/asciiui"

func main() {
	theme := asciiui.NewPrettyTheme()
	ui := asciiui.NewUI(24, 13, theme)

	ui.AddComponent(asciiui.NewButton(3, 1, "Hi, I miss u", asciiui.Padding{Width: 1, Height: 0}))
	ui.AddComponent(asciiui.NewText(4, 4, "Do u love me ?\nPlease tell..."))
	ui.AddComponent(asciiui.NewButton(3, 6, "No", asciiui.Padding{Width: 1, Height: 0}))
	ui.AddComponent(asciiui.NewButton(12, 6, "Yes", asciiui.Padding{Width: 1, Height: 0}))
	ui.AddComponent(asciiui.NewNumberedList(3, 9, []string{"Let's Travel", "Back to home", "Have dinner"}))

	fmt.Println(ui.Render())
}
