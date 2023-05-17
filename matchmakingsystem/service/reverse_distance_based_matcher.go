package service

import (
	"matchmakingsystem/entity"
)

type ReverseDistanceBasedMatcher struct {
	DistanceBasedMatcher
}

func NewReverseDistanceBasedMatcher() *ReverseDistanceBasedMatcher {
	m := &ReverseDistanceBasedMatcher{}
	m.UnimplementedStrategy = interface{}(m).(UnimplementedStrategy)
	return m
}

func (m *ReverseDistanceBasedMatcher) SortingStrategy(hmwa []entity.HowMatchWeAre) entity.Individual {
	var farthest entity.Individual
	maxDistance := float64(0)
	for _, h := range hmwa {
		if h.MatchLevel.(float64) > maxDistance {
			maxDistance = h.MatchLevel.(float64)
			farthest = h.Individual
		}
	}

	return farthest
}
