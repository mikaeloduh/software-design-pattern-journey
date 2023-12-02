package service

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRealDatabaseProtectionProxy_GetEmployeeById(t *testing.T) {
	db := NewRealDatabaseProtectionProxy()
	db.Init()

	t.Run("test GetEmployeeById password correct", func(t *testing.T) {
		_ = os.Setenv("PASSWORD", "1qaz2wsx")

		got, err := db.GetEmployeeById(1)

		assert.NoError(t, err)
		assert.Equal(t, 1, got.Id())
		assert.Equal(t, "waterball", got.Name())
	})

	t.Run("test GetEmployeeById password incorrect", func(t *testing.T) {
		_ = os.Setenv("PASSWORD", "incorrect_password")

		_, err := db.GetEmployeeById(1)

		assert.Error(t, err)
	})

	t.Run("test given id should return correct employee info and subordinates", func(t *testing.T) {
		_ = os.Setenv("PASSWORD", "1qaz2wsx")

		got, err := db.GetEmployeeById(2)

		assert.NoError(t, err)
		assert.Equal(t, 2, got.Id())

		assert.ElementsMatch(t, []Employee{
			NewRealEmployee(1, "waterball", 25, nil),
			NewRealEmployee(3, "fong", 7, []int{1}),
		}, got.Subordinates())
	})
}
