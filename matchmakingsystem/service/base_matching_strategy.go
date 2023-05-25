package service

import (
	"fmt"
	"matchmakingsystem/entity"
)

type UnimplementedMethod interface {
	Search(me entity.Individual, other entity.Individual) interface{}
	Sort([]entity.HowMatchWeAre) entity.Individual
}

type BaseMatchingStrategy struct {
	UnimplementedMethod
}

func (m *BaseMatchingStrategy) Match(me entity.Individual, others []entity.Individual) (entity.Individual, error) {
	if len(others) == 0 {
		return entity.Individual{}, fmt.Errorf("no individuals to match")
	}

	howMatchWeAres := m.calculateMatchLevel(me, others)
	bestMatch := m.UnimplementedMethod.Sort(howMatchWeAres)

	return bestMatch, nil
}

func (m *BaseMatchingStrategy) calculateMatchLevel(me entity.Individual, others []entity.Individual) []entity.HowMatchWeAre {
	var hmwr []entity.HowMatchWeAre
	for _, o := range others {
		if o.Id == me.Id {
			continue
		}

		distance := m.UnimplementedMethod.Search(me, o)
		hmwr = append(hmwr, entity.HowMatchWeAre{Individual: o, MatchLevel: distance})
	}

	return hmwr
}

func (m *BaseMatchingStrategy) Search(me entity.Individual, other entity.Individual) interface{} {
	panic("method Search not implemented")
}

func (m *BaseMatchingStrategy) Sort([]entity.HowMatchWeAre) entity.Individual {
	panic("method Sort not implemented")
}
