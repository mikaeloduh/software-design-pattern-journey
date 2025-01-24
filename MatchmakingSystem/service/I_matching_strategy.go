package service

import "matchmakingsystem/entity"

type IMatchingStrategy interface {
	Match(me entity.Individual, others []entity.Individual) (entity.Individual, error)
}
