package main

import (
	"fmt"
	"matchmakingsystem/entity"
	"matchmakingsystem/service"
)

func main() {
	var p1 = entity.Individual{
		Id:     1,
		Gender: entity.Male,
		Age:    10,
		Intro:  "Hello Intro",
		Habits: []string{"baseball", "cook", "sleep"},
		Coord: entity.Coord{
			X: 10,
			Y: 10,
		},
	}

	var p2 = entity.Individual{
		Id:     2,
		Gender: entity.Female,
		Age:    20,
		Intro:  "Hi there",
		Habits: []string{"music", "sleep", "travel"},
		Coord: entity.Coord{
			X: 5,
			Y: 5,
		},
	}

	var p3 = entity.Individual{
		Id:     3,
		Gender: entity.Other,
		Age:    30,
		Intro:  "Hey",
		Habits: []string{"baseball", "sports", "reading", "sleep"},
		Coord: entity.Coord{
			X: 15,
			Y: 15,
		},
	}

	pps := []entity.Individual{p1, p2, p3}
	distanceBasedMatcher := service.NewDistanceBasedMatcher()
	habitBasedMatcher := service.NewHabitBasedMatcher()

	closest1, _ := distanceBasedMatcher.Match(p1, pps)
	fmt.Printf("%d's best match is %d\n", p1.Id, closest1.Id)
	closest2, _ := habitBasedMatcher.Match(p2, pps)
	fmt.Printf("%d's best match is %d\n", p2.Id, closest2.Id)
	closest3, _ := habitBasedMatcher.Match(p3, pps)
	fmt.Printf("%d's best match is %d\n", p3.Id, closest3.Id)
}
