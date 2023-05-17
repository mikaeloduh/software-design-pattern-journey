package service

import (
	"matchmakingsystem/entity"
	"math"
)

type ReverseHabitBasedMatcher struct {
	HabitBasedMatcher
}

func NewReverseHabitBasedMatcher() *ReverseHabitBasedMatcher {
	m := &ReverseHabitBasedMatcher{}
	m.UnimplementedStrategy = interface{}(m).(UnimplementedStrategy)
	return m
}

func (m *ReverseHabitBasedMatcher) SortingStrategy(hmwa []entity.HowMatchWeAre) entity.Individual {
	var lowestSimilarity entity.Individual
	minIntersection := math.MaxInt
	for _, i := range hmwa {
		if i.MatchLevel.(int) < minIntersection {
			minIntersection = i.MatchLevel.(int)
			lowestSimilarity = i.Individual
		}
	}

	return lowestSimilarity
}
