package main

import (
	"matchmakingsystem/entity"
	"matchmakingsystem/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchDistanceBased(t *testing.T) {
	p1 := entity.Individual{
		Id:     1,
		Gender: entity.Male,
		Age:    10,
		Intro:  "Hello Intro",
		Habits: []string{"baseball", "cook", "sleep"},
		Coord: entity.Coord{
			X: 10,
			Y: 10,
		},
	}
	p2 := entity.Individual{
		Id:     2,
		Gender: entity.Female,
		Age:    20,
		Intro:  "Hi there",
		Habits: []string{"cook", "sleep", "dance"},
		Coord: entity.Coord{
			X: 5,
			Y: 5,
		},
	}
	p3 := entity.Individual{
		Id:     3,
		Gender: entity.Male,
		Age:    30,
		Intro:  "Nice to meet you",
		Habits: []string{"baseball", "sleep", "read"},
		Coord: entity.Coord{
			X: 15,
			Y: 15,
		},
	}

	m := service.NewDistanceBasedMatcher()

	// Test matching with empty individuals slice
	_, err := m.Match(p1, []entity.Individual{})
	assert.Error(t, err)

	// Test matching with one Individual
	ind, err := m.Match(p1, []entity.Individual{p2})
	assert.NoError(t, err)
	assert.Equal(t, p2, ind)

	// Test matching with multiple individuals
	ind, err = m.Match(p1, []entity.Individual{p2, p3})
	assert.NoError(t, err)
	assert.Equal(t, p2, ind)

	// Test case 1: Matchmaker matches p1 and p2
	result1, err1 := m.Match(p1, []entity.Individual{p2, p3})
	assert.Nil(t, err1)
	assert.Equal(t, p2, result1)

	// Test case 2: Matchmaker matches p1 and p3
	result2, err2 := m.Match(p1, []entity.Individual{p3, p2})
	assert.Nil(t, err2)
	assert.Equal(t, p3, result2)

	// Test case 3: Matchmaker returns an error if individuals slice is empty
	result3, err3 := m.Match(p1, []entity.Individual{})
	assert.NotNil(t, err3)
	assert.EqualError(t, err3, "no individuals to match")
	assert.Equal(t, entity.Individual{}, result3)

	m2 := service.NewHabitBasedMatcher()
	// Test case 4: HabitBased
	result4, err := m2.Match(p1, []entity.Individual{p2, p3})
	assert.NoError(t, err)
	assert.NotNil(t, result4)
	assert.Equal(t, p2, result4)
}
