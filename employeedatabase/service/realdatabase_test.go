package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRealDatabase_GetEmployeeById(t *testing.T) {
	db := NewRealDatabase("data.txt")
	err := db.Init()

	assert.NoError(t, err)

	t.Run("test given id should return correct employee info", func(t *testing.T) {

		got, err := db.GetEmployeeById(1)

		assert.NoError(t, err)
		assert.Equal(t, 1, got.id)
		assert.Equal(t, "waterball", got.name)
	})

	t.Run("test given id should return correct employee info (cont)", func(t *testing.T) {

		got, err := db.GetEmployeeById(5)

		assert.NoError(t, err)
		assert.Equal(t, 5, got.id)
		assert.Equal(t, "peterchen", got.name)
	})
}
