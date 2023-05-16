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

type HowMatchWeAre struct {
	individual Individual
	matchLevel interface{}
}

type UnimplementedStrategy interface {
	CalculateStrategy(me Individual, other Individual) interface{}
	SortingStrategy([]HowMatchWeAre) Individual
}

type BaseMatcher struct {
	UnimplementedStrategy
}

func (m *BaseMatcher) Match(me Individual, others []Individual) (Individual, error) {
	if len(others) == 0 {
		return Individual{}, fmt.Errorf("no individuals to match")
	}

	howMatchWeAres := m.calculateMatchLevel(me, others)
	bestMatch := m.UnimplementedStrategy.SortingStrategy(howMatchWeAres)

	return bestMatch, nil
}

func (m *BaseMatcher) calculateMatchLevel(me Individual, others []Individual) []HowMatchWeAre {
	var hmwa []HowMatchWeAre
	for _, o := range others {
		if o.id == me.id {
			continue
		}

		distance := m.UnimplementedStrategy.CalculateStrategy(me, o)
		hmwa = append(hmwa, HowMatchWeAre{o, distance})
	}

	return hmwa
}

type DistanceBasedMatcher struct {
	BaseMatcher
}

func NewDistanceBasedMatcher() *DistanceBasedMatcher {
	m := &DistanceBasedMatcher{}
	m.UnimplementedStrategy = interface{}(m).(UnimplementedStrategy)
	return m
}

func (m *DistanceBasedMatcher) CalculateStrategy(me Individual, other Individual) interface{} {
	distance := other.coord.distanceTo(me.coord)

	return distance
}

func (m *DistanceBasedMatcher) SortingStrategy(hmwa []HowMatchWeAre) Individual {
	var closest Individual
	minDistance := math.MaxFloat64
	for _, h := range hmwa {
		if h.matchLevel.(float64) < minDistance {
			minDistance = h.matchLevel.(float64)
			closest = h.individual
		}
	}

	return closest
}

type HabitBasedMatcher struct {
	UnimplementedStrategy
}

func NewHabitBasedMatcher() *HabitBasedMatcher {
	m := &HabitBasedMatcher{}
	m.UnimplementedStrategy = interface{}(m).(UnimplementedStrategy)
	return m
}

func (m *HabitBasedMatcher) Match(me Individual, others []Individual) (Individual, error) {
	if len(others) == 0 {
		return Individual{}, fmt.Errorf("no individuals to match")
	}

	hmwa := m.calculateMatchLevel(me, others)
	bestMatch := m.SortingStrategy(hmwa)

	return bestMatch, nil
}

func (m *HabitBasedMatcher) calculateMatchLevel(me Individual, others []Individual) []HowMatchWeAre {
	var hmwa []HowMatchWeAre
	for _, i := range others {
		if i.id == me.id {
			continue
		}

		intersection := m.CalculateStrategy(me, i)
		hmwa = append(hmwa, HowMatchWeAre{i, intersection})
	}
	return hmwa
}

func (m *HabitBasedMatcher) CalculateStrategy(target Individual, i Individual) interface{} {
	count := 0
	for _, x := range target.habits {
		for _, y := range i.habits {
			if x == y {
				count++
				break
			}
		}
	}
	return count
}
func (m *HabitBasedMatcher) SortingStrategy(hmwa []HowMatchWeAre) Individual {
	var highestSimilarity Individual
	maxIntersection := 0
	for _, i := range hmwa {
		if i.matchLevel.(int) > maxIntersection {
			maxIntersection = i.matchLevel.(int)
			highestSimilarity = i.individual
		}
	}

	return highestSimilarity
}

func main() {
	var p1 = Individual{
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

	var p2 = Individual{
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

	var p3 = Individual{
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

	pps := []Individual{p1, p2, p3}
	distanceBasedMatcher := NewDistanceBasedMatcher()
	habitBasedMatcher := NewHabitBasedMatcher()

	closest1, _ := distanceBasedMatcher.Match(p1, pps)
	fmt.Println(closest1.id)
	closest2, _ := habitBasedMatcher.Match(p2, pps)
	fmt.Println(closest2.id)
	closest3, _ := habitBasedMatcher.Match(p3, pps)
	fmt.Println(closest3.id)
}
