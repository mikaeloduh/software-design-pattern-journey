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
			X: 0,
			Y: 0,
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
	p4 := entity.Individual{
		Id:     4,
		Gender: entity.Male,
		Age:    30,
		Intro:  "Nice to meet you",
		Habits: []string{"drink", "smoke", "cook"},
		Coord: entity.Coord{
			X: 5,
			Y: 5,
		},
	}

	// Testing Distance-based strategy
	m := service.NewDistanceBasedMatcher()

	t.Run("Should return an error if individuals slice is empty", func(t *testing.T) {
		// Test case: Should return an error if individuals slice is empty
		result, err := m.Match(p1, []entity.Individual{})
		assert.NotNil(t, err)
		assert.EqualError(t, err, "no individuals to match")
		assert.Equal(t, entity.Individual{}, result)
	})

	t.Run("Test matching with one Individual", func(t *testing.T) {
		// Test case: Test matching with one Individual
		ind, err := m.Match(p1, []entity.Individual{p2})
		assert.NoError(t, err)
		assert.Equal(t, p2, ind)
	})

	t.Run("Should return the closest match (p2)", func(t *testing.T) {
		// Test case: Should return the closest match (p2)
		result1, err1 := m.Match(p1, []entity.Individual{p2, p3})
		assert.NoError(t, err1)
		assert.Equal(t, p2, result1)
	})

	t.Run("Should return the closest match (p2)", func(t *testing.T) {
		// Test case 2: Should return the closest match (p2)
		result2, err2 := m.Match(p1, []entity.Individual{p3, p2})
		assert.NoError(t, err2)
		assert.Equal(t, p2, result2)
	})

	// Testing Habit-based strategy
	m2 := service.NewHabitBasedMatcher()

	t.Run("best match should return p2", func(t *testing.T) {
		// Test case: best match should return p2
		result4, err := m2.Match(p1, []entity.Individual{p2, p3})
		assert.NoError(t, err)
		assert.NotNil(t, result4)
		assert.Equal(t, p2, result4)
	})

	// Testing Reverse Distance-based strategy
	m3 := service.NewReverseDistanceBasedMatcher()

	t.Run("Should return the farthest match (p3)", func(t *testing.T) {
		// Test case: Should return the farthest match (p3)
		result5, err1 := m3.Match(p1, []entity.Individual{p2, p3})
		assert.NoError(t, err1)
		assert.Equal(t, p3, result5)
	})

	// Testing Reverse Habit-based strategy
	m4 := service.NewReverseHabitBasedMatcher()

	t.Run("Should return the lowestSimilarity match (p4)", func(t *testing.T) {
		// Test case: Should return the lowestSimilarity match (p4)
		result6, err6 := m4.Match(p1, []entity.Individual{p2, p3, p4})
		assert.NoError(t, err6)
		assert.Equal(t, p4, result6)
	})
}
