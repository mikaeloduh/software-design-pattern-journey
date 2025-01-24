package service

import (
	"fmt"
	"matchmakingsystem/entity"
)

type UnimplementedMethod interface {
	Count(me entity.Individual, other entity.Individual) interface{}
	Find([]entity.HowMatchWeAre) entity.Individual
}

type BaseMatchingStrategy struct {
	UnimplementedMethod
}

func (s *BaseMatchingStrategy) Match(me entity.Individual, others []entity.Individual) (entity.Individual, error) {
	if len(others) == 0 {
		return entity.Individual{}, fmt.Errorf("no individuals to match")
	}

	howMatchWeAres := s.calculateMatchLevel(me, others)
	bestMatch := s.UnimplementedMethod.Find(howMatchWeAres)

	return bestMatch, nil
}

func (s *BaseMatchingStrategy) calculateMatchLevel(me entity.Individual, others []entity.Individual) []entity.HowMatchWeAre {
	var hmwr []entity.HowMatchWeAre
	for _, o := range others {
		if o.Id == me.Id {
			continue
		}

		distance := s.UnimplementedMethod.Count(me, o)
		hmwr = append(hmwr, entity.HowMatchWeAre{Individual: o, MatchLevel: distance})
	}

	return hmwr
}

func (s *BaseMatchingStrategy) Count(me entity.Individual, other entity.Individual) interface{} {
	panic("method Count not implemented")
}

func (s *BaseMatchingStrategy) Find([]entity.HowMatchWeAre) entity.Individual {
	panic("method Find not implemented")
}
