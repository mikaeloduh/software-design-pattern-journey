package service

import "matchmakingsystem/entity"

type HabitBasedMatchingStrategy struct {
	BaseMatchingStrategy
}

func NewHabitBasedMatchingStrategy() *HabitBasedMatchingStrategy {
	h := &HabitBasedMatchingStrategy{}
	h.UnimplementedMethod = interface{}(h).(UnimplementedMethod)
	return h
}

func (s *HabitBasedMatchingStrategy) Count(me entity.Individual, other entity.Individual) interface{} {
	return me.Habits.CountIntersection(other.Habits)
}

func (s *HabitBasedMatchingStrategy) Find(hmwr []entity.HowMatchWeAre) entity.Individual {
	var highestSimilarity entity.Individual
	maxIntersection := 0
	for _, i := range hmwr {
		if i.MatchLevel.(int) > maxIntersection {
			maxIntersection = i.MatchLevel.(int)
			highestSimilarity = i.Individual
		}
	}

	return highestSimilarity
}
