package service

import (
	"matchmakingsystem/entity"
	"math"
)

type DistanceBasedMatchingStrategy struct {
	BaseMatchingStrategy
}

func NewDistanceBasedMatchingStrategy() *DistanceBasedMatchingStrategy {
	d := &DistanceBasedMatchingStrategy{}
	d.UnimplementedMethod = interface{}(d).(UnimplementedMethod)
	return d
}

func (s *DistanceBasedMatchingStrategy) Count(me entity.Individual, other entity.Individual) interface{} {
	return other.Coord.DistanceTo(me.Coord)
}

func (s *DistanceBasedMatchingStrategy) Find(hmwr []entity.HowMatchWeAre) entity.Individual {
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
