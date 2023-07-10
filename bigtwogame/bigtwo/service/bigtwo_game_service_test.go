package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBigTwo(t *testing.T) {
	t.Run("Test 1+1=2", func(t *testing.T) {
		assert.Equal(t, 2, 1+1)
	})
}
