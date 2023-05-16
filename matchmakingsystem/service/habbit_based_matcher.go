package service

import "matchmakingsystem/entity"

type HabitBasedMatcher struct {
	BaseMatcher
}

func NewHabitBasedMatcher() *HabitBasedMatcher {
	m := &HabitBasedMatcher{}
	m.UnimplementedStrategy = interface{}(m).(UnimplementedStrategy)
	return m
}

func (m *HabitBasedMatcher) CalculateStrategy(me entity.Individual, other entity.Individual) interface{} {
	count := 0
	for _, x := range me.Habits {
		for _, y := range other.Habits {
			if x == y {
				count++
				break
			}
		}
	}
	return count
}
func (m *HabitBasedMatcher) SortingStrategy(hmwa []entity.HowMatchWeAre) entity.Individual {
	var highestSimilarity entity.Individual
	maxIntersection := 0
	for _, i := range hmwa {
		if i.MatchLevel.(int) > maxIntersection {
			maxIntersection = i.MatchLevel.(int)
			highestSimilarity = i.Individual
		}
	}

	return highestSimilarity
}
