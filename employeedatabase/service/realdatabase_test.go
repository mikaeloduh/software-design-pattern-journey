package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRealDatabase_GetEmployeeById(t *testing.T) {
	t.Run("test given id should return correct employee info", func(t *testing.T) {

		db := NewRealDatabase()
		got, _ := db.GetEmployeeById(1)

		assert.Equal(t, 1, got.Id)
		assert.Equal(t, "waterball", got.Name)
	})
}
