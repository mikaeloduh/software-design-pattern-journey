package service

import (
	"matchmakingsystem/entity"
)

type ReverseDistanceBasedMatchingStrategy struct {
	DistanceBasedMatchingStrategy
}

func NewReverseDistanceBasedMatchingStrategy() *ReverseDistanceBasedMatchingStrategy {
	m := &ReverseDistanceBasedMatchingStrategy{}
	m.UnimplementedMethod = interface{}(m).(UnimplementedMethod)
	return m
}

func (m *ReverseDistanceBasedMatchingStrategy) Sort(hmwr []entity.HowMatchWeAre) entity.Individual {
	var farthest entity.Individual
	maxDistance := float64(0)
	for _, h := range hmwr {
		if h.MatchLevel.(float64) > maxDistance {
			maxDistance = h.MatchLevel.(float64)
			farthest = h.Individual
		}
	}

	return farthest
}
