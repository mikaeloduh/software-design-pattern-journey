package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	tests := []struct {
		name string
		want *Game
	}{
		{
			"TestHello",
			&Game{},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := NewGame()

			assert.Equal(t, tc.want, got)
		})
	}
}
