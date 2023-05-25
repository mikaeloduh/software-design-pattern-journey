package service

import (
	"matchmakingsystem/entity"
	"math"
)

type DistanceBasedMatchingStrategy struct {
	BaseMatchingStrategy
}

func NewDistanceBasedMatchingStrategy() *DistanceBasedMatchingStrategy {
	m := &DistanceBasedMatchingStrategy{}
	m.UnimplementedMethod = interface{}(m).(UnimplementedMethod)
	return m
}

func (m *DistanceBasedMatchingStrategy) Search(me entity.Individual, other entity.Individual) interface{} {
	distance := other.Coord.DistanceTo(me.Coord)

	return distance
}

func (m *DistanceBasedMatchingStrategy) Sort(hmwr []entity.HowMatchWeAre) entity.Individual {
	var closest entity.Individual
	minDistance := math.MaxFloat64
	for _, h := range hmwr {
		if h.MatchLevel.(float64) < minDistance {
			minDistance = h.MatchLevel.(float64)
			closest = h.Individual
		}
	}

	return closest
}
