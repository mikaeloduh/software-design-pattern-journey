package service

import (
	"fmt"
	"matchmakingsystem/entity"
	"matchmakingsystem/samples"
)

type MatchmakingSystem struct {
	matcher IMatcher
}

func NewMatchmakingSystem(matcher IMatcher) *MatchmakingSystem {
	return &MatchmakingSystem{matcher: matcher}
}

func (s MatchmakingSystem) Run() {
	individuals := []entity.Individual{samples.P1, samples.P2, samples.P3}

	for _, p := range individuals {
		bestMatch, _ := s.matcher.Match(p, individuals)
		fmt.Printf("%d's best match is %d\n", p.Id, bestMatch.Id)
	}
}
