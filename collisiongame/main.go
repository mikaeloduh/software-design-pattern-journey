package main

func main() {
	println("hello world")

	handler := NewHeroHeroHandler(NewHeroWaterHandler(NewHeroFireHandler(NewWaterHeroHandler(NewWaterWaterHandler(NewWaterFireHandler(NewFireHeroHandler(NewFireWaterHandler(NewFireFireHandler(nil)))))))))
	world := NewWorld(handler)

	for {
		x1, x2 := inputCoord()
		world.Move(x1, x2)
	}
}
