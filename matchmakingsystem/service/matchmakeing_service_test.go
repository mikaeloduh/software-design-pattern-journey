package service

import (
	"reflect"
	"testing"
)

func TestNewMatchmaking(t *testing.T) {
	tests := []struct {
		name string
		want *Matchmaking
	}{
		// TODO: Add test cases.
		{
			"Hello world",
			&Matchmaking{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatchmaking(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatchmaking() = %v, want %v", got, tt.want)
			}
		})
	}
}
