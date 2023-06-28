package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	world := World{}

	assert.IsType(t, World{}, world)

	world.Init()

	assert.Equal(t, 30, len(world.coord))
}
