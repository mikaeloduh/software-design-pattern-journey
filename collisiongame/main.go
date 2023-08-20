package main

import (
	"collisiongame/commons"
	"collisiongame/entity"
)

func main() {
	println("hello world")

	handler := entity.NewHeroHeroHandler(
		entity.NewHeroWaterHandler(
			entity.NewHeroFireHandler(
				entity.NewWaterHeroHandler(
					entity.NewWaterWaterHandler(
						entity.NewWaterFireHandler(
							entity.NewFireHeroHandler(
								entity.NewFireWaterHandler(
									entity.NewFireFireHandler(nil)))))))))
	world := entity.NewWorld(handler)

	for {
		x1, x2 := commons.InputCoord()
		world.Move(x1, x2)
	}
}
