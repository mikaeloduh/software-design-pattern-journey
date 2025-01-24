package service

import (
	"fmt"
	"matchmakingsystem/entity"
)

type MatchmakingSystem struct {
	matcher     IMatchingStrategy
	individuals []entity.Individual
}

func NewMatchmakingSystem(matcher IMatchingStrategy, i []entity.Individual) *MatchmakingSystem {
	return &MatchmakingSystem{matcher: matcher, individuals: i}
}

func (s MatchmakingSystem) Run() {
	for _, i := range s.individuals {
		bestMatch, _ := s.matcher.Match(i, s.individuals)
		fmt.Printf("%d's best match is %d\n", i.Id, bestMatch.Id)
	}
}
