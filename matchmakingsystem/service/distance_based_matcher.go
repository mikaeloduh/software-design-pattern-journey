package service

import (
	"matchmakingsystem/entity"
	"math"
)

type DistanceBasedMatcher struct {
	BaseMatcher
}

func NewDistanceBasedMatcher() *DistanceBasedMatcher {
	m := &DistanceBasedMatcher{}
	m.UnimplementedStrategy = interface{}(m).(UnimplementedStrategy)
	return m
}

func (m *DistanceBasedMatcher) CalculateStrategy(me entity.Individual, other entity.Individual) interface{} {
	distance := other.Coord.DistanceTo(me.Coord)

	return distance
}

func (m *DistanceBasedMatcher) SortingStrategy(hmwa []entity.HowMatchWeAre) entity.Individual {
	var closest entity.Individual
	minDistance := math.MaxFloat64
	for _, h := range hmwa {
		if h.MatchLevel.(float64) < minDistance {
			minDistance = h.MatchLevel.(float64)
			closest = h.Individual
		}
	}

	return closest
}
