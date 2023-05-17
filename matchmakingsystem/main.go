package main

import (
	"matchmakingsystem/service"
)

func main() {
	service.NewMatchmakingSystem(service.NewDistanceBasedMatcher()).Run()
	service.NewMatchmakingSystem(service.NewHabitBasedMatcher()).Run()
	service.NewMatchmakingSystem(service.NewReverseDistanceBasedMatcher()).Run()
	service.NewMatchmakingSystem(service.NewReverseHabitBasedMatcher()).Run()
}
