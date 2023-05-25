package main

import (
	"matchmakingsystem/entity"
	"matchmakingsystem/samples"
	"matchmakingsystem/service"
)

func main() {
	individuals := []entity.Individual{samples.P1, samples.P2, samples.P3}
	service.NewMatchmakingSystem(service.NewDistanceBasedMatchingStrategy(), individuals).Run()
	service.NewMatchmakingSystem(service.NewHabitBasedMatchingStrategy(), individuals).Run()
	service.NewMatchmakingSystem(service.NewReverseDistanceBasedMatchingStrategy(), individuals).Run()
	service.NewMatchmakingSystem(service.NewReverseHabitBasedMatchingStrategy(), individuals).Run()
}
