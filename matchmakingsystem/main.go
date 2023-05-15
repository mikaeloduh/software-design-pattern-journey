package main

import (
	"fmt"
	"math"
)

type Gender int

const (
	Male Gender = iota
	Female
	Other
)

type Coord struct {
	x float64
	y float64
}

func (c Coord) distanceTo(other Coord) float64 {
	return math.Sqrt((c.x-other.x)*(c.x-other.x) + (c.y-other.y)*(c.y-other.y))
}

type Individual struct {
	id     int
	gender Gender
	age    int
	intro  string
	habits []string
	coord  Coord
}

type Matchmaker struct {
}

func (m *Matchmaker) MatchDistanceBased(target Individual, individuals []Individual) (Individual, error) {
	if len(individuals) == 0 {
		return Individual{}, fmt.Errorf("no individuals to match")
	}

	var closest Individual
	minDistance := math.MaxFloat64

	for _, i := range individuals {
		if i.id == target.id {
			continue
		}

		distance := i.coord.distanceTo(target.coord)

		if distance < minDistance {
			minDistance = distance
			closest = i
		}
	}

	return closest, nil
}

func (m *Matchmaker) MatchHabitBased(target Individual, individuals []Individual) (Individual, error) {
	if len(individuals) == 0 {
		return Individual{}, fmt.Errorf("no individuals to match")
	}

	var bestMatch Individual
	maxIntersection := 0

	for _, i := range individuals {
		if i.id == target.id {
			continue
		}

		intersection := countIntersection(target.habits, i.habits)
		if intersection > maxIntersection {
			maxIntersection = intersection
			bestMatch = i
		}
	}

	return bestMatch, nil
}

func countIntersection(a, b []string) int {
	count := 0
	for _, x := range a {
		for _, y := range b {
			if x == y {
				count++
				break
			}
		}
	}
	return count
}

func main() {
	var p1 Individual = Individual{
		id:     1,
		gender: Male,
		age:    10,
		intro:  "Hello intro",
		habits: []string{"baseball", "cook", "sleep"},
		coord: Coord{
			x: 10,
			y: 10,
		},
	}

	var p2 Individual = Individual{
		id:     2,
		gender: Female,
		age:    20,
		intro:  "Hi there",
		habits: []string{"music", "sleep", "travel"},
		coord: Coord{
			x: 5,
			y: 5,
		},
	}

	var p3 Individual = Individual{
		id:     3,
		gender: Other,
		age:    30,
		intro:  "Hey",
		habits: []string{"baseball", "sports", "reading", "sleep"},
		coord: Coord{
			x: 15,
			y: 15,
		},
	}

	matchmaker := Matchmaker{}
	pps := []Individual{p1, p2, p3}

	closest1, _ := matchmaker.MatchHabitBased(p1, pps)
	fmt.Println(closest1.id)
	closest2, _ := matchmaker.MatchHabitBased(p2, pps)
	fmt.Println(closest2.id)
	closest3, _ := matchmaker.MatchHabitBased(p3, pps)
	fmt.Println(closest3.id)
}
