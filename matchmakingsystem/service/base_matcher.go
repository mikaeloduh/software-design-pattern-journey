package service

import (
	"fmt"
	"matchmakingsystem/entity"
)

type UnimplementedStrategy interface {
	CalculateStrategy(me entity.Individual, other entity.Individual) interface{}
	SortingStrategy([]entity.HowMatchWeAre) entity.Individual
}

type BaseMatcher struct {
	UnimplementedStrategy
}

func (m *BaseMatcher) Match(me entity.Individual, others []entity.Individual) (entity.Individual, error) {
	if len(others) == 0 {
		return entity.Individual{}, fmt.Errorf("no individuals to match")
	}

	howMatchWeAres := m.calculateMatchLevel(me, others)
	bestMatch := m.UnimplementedStrategy.SortingStrategy(howMatchWeAres)

	return bestMatch, nil
}

func (m *BaseMatcher) calculateMatchLevel(me entity.Individual, others []entity.Individual) []entity.HowMatchWeAre {
	var hmwa []entity.HowMatchWeAre
	for _, o := range others {
		if o.Id == me.Id {
			continue
		}

		distance := m.UnimplementedStrategy.CalculateStrategy(me, o)
		hmwa = append(hmwa, entity.HowMatchWeAre{Individual: o, MatchLevel: distance})
	}

	return hmwa
}

func (m *BaseMatcher) CalculateStrategy(me entity.Individual, other entity.Individual) interface{} {
	panic("method CalculateStrategy not implemented")
}

func (m *BaseMatcher) SortingStrategy([]entity.HowMatchWeAre) entity.Individual {
	panic("method SortingStrategy not implemented")
}
