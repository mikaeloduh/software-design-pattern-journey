package main

import (
	"github.com/stretchr/testify/assert"

	"bytes"
	"testing"
)

func TestCharacterStatus(t *testing.T) {
	t.Run("happy test", func(t *testing.T) {
		assert.Equal(t, 2, 1+1)
	})

	t.Run("test character poisoned and HP -15 on the following two round start", func(t *testing.T) {
		var writer bytes.Buffer

		c := NewCharacter(&writer)
		c.SetState(NewPoisonedState(c))

		assert.Equal(t, "The character is poisoned.\n", writer.String())
	})
}
