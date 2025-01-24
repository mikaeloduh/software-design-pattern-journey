package service

import (
	"matchmakingsystem/entity"
	"math"
)

type ReverseHabitBasedMatchingStrategy struct {
	HabitBasedMatchingStrategy
}

func NewReverseHabitBasedMatchingStrategy() *ReverseHabitBasedMatchingStrategy {
	m := &ReverseHabitBasedMatchingStrategy{}
	m.UnimplementedMethod = interface{}(m).(UnimplementedMethod)
	return m
}

func (m *ReverseHabitBasedMatchingStrategy) Find(hmwr []entity.HowMatchWeAre) entity.Individual {
	var lowestSimilarity entity.Individual
	minIntersection := math.MaxInt
	for _, i := range hmwr {
		if i.MatchLevel.(int) < minIntersection {
			minIntersection = i.MatchLevel.(int)
			lowestSimilarity = i.Individual
		}
	}

	return lowestSimilarity
}
