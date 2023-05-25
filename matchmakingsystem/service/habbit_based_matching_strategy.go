package service

import "matchmakingsystem/entity"

type HabitBasedMatchingStrategy struct {
	BaseMatchingStrategy
}

func NewHabitBasedMatchingStrategy() *HabitBasedMatchingStrategy {
	m := &HabitBasedMatchingStrategy{}
	m.UnimplementedMethod = interface{}(m).(UnimplementedMethod)
	return m
}

func (m *HabitBasedMatchingStrategy) Search(me entity.Individual, other entity.Individual) interface{} {
	return me.Habits.CountIntersection(other.Habits)
}
func (m *HabitBasedMatchingStrategy) Sort(hmwr []entity.HowMatchWeAre) entity.Individual {
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
