package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMatchDistanceBased(t *testing.T) {
	p1 := Individual{
		id:     1,
		gender: Male,
		age:    10,
		intro:  "Hello intro",
		habits: []string{"baseball", "cook", "sleep"},
		coord: Coord{
			x: 10,
			y: 10,
		},
	}
	p2 := Individual{
		id:     2,
		gender: Female,
		age:    20,
		intro:  "Hi there",
		habits: []string{"cook", "sleep", "dance"},
		coord: Coord{
			x: 5,
			y: 5,
		},
	}
	p3 := Individual{
		id:     3,
		gender: Male,
		age:    30,
		intro:  "Nice to meet you",
		habits: []string{"baseball", "sleep", "read"},
		coord: Coord{
			x: 15,
			y: 15,
		},
	}

	m := NewDistanceBasedMatcher()

	// Test matching with empty individuals slice
	_, err := m.Match(p1, []Individual{})
	assert.Error(t, err)

	// Test matching with one individual
	ind, err := m.Match(p1, []Individual{p2})
	assert.NoError(t, err)
	assert.Equal(t, p2, ind)

	// Test matching with multiple individuals
	ind, err = m.Match(p1, []Individual{p2, p3})
	assert.NoError(t, err)
	assert.Equal(t, p2, ind)

	// Test case 1: Matchmaker matches p1 and p2
	result1, err1 := m.Match(p1, []Individual{p2, p3})
	assert.Nil(t, err1)
	assert.Equal(t, p2, result1)

	// Test case 2: Matchmaker matches p1 and p3
	result2, err2 := m.Match(p1, []Individual{p3, p2})
	assert.Nil(t, err2)
	assert.Equal(t, p3, result2)

	// Test case 3: Matchmaker returns an error if individuals slice is empty
	result3, err3 := m.Match(p1, []Individual{})
	assert.NotNil(t, err3)
	assert.EqualError(t, err3, "no individuals to match")
	assert.Equal(t, Individual{}, result3)

	m2 := NewHabitBasedMatcher()
	// Test case 4: HabitBased
	result4, err := m2.Match(p1, []Individual{p2, p3})
	assert.NoError(t, err)
	assert.NotNil(t, result4)
	assert.Equal(t, p2, result4)
}
