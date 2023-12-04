package service

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestRealDatabaseProtectionProxy_GetEmployeeById(t *testing.T) {
	db := FakeNewRealDatabaseProtectionProxy()

	t.Run("test GetEmployeeById password correct", func(t *testing.T) {
		_ = os.Setenv("PASSWORD", "1qaz2wsx")

		_, err := db.GetEmployeeById(1)

		assert.NoError(t, err)
	})

	t.Run("test GetEmployeeById password incorrect", func(t *testing.T) {
		_ = os.Setenv("PASSWORD", "incorrect_password")

		_, err := db.GetEmployeeById(1)

		assert.Error(t, err)
	})
}

func FakeNewRealDatabaseProtectionProxy() *RealDatabaseProtectionProxy {
	return NewRealDatabaseProtectionProxy(&RealDatabase{})
}
