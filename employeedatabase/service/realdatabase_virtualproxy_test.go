package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRealDatabaseVirtualProxy_GetEmployeeById(t *testing.T) {
	db := NewRealDatabaseVirtualProxy()

	t.Run("test given id should return correct employee info and subordinates", func(t *testing.T) {

		got, err := db.GetEmployeeById(2)

		assert.NoError(t, err)
		assert.Equal(t, 2, got.Id())

		assert.ElementsMatch(t, []IEmployee{
			NewRealEmployee(1, "waterball", 25, nil),
			NewRealEmployee(3, "fong", 7, []int{1}),
		}, got.Subordinates())
	})
}
