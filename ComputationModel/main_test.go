package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"computationmodel/mod"
)

func Test_Main(t *testing.T) {
	models := mod.NewModels()

	t.Run("Creating Scaling model should return correct transform matrix", func(t *testing.T) {
		scalingModel, err := models.CreateModel("Scaling")
		assert.NoError(t, err)

		array := testArray(1.0, 1000)

		output, err2 := scalingModel.LinearTransformation(array)
		assert.NoError(t, err2)

		for i := range output {
			assert.Equal(t, 2.0, output[i])
		}
		assert.Implements(t, (*mod.IModel)(nil), scalingModel)
	})

	var scalingModel mod.IModel

	t.Run("Validating array length must equal the model's row size", func(t *testing.T) {
		var err error
		scalingModel, err = models.CreateModel("Scaling")
		assert.NoError(t, err)
		array := testArray(1.0, 999)

		_, err2 := scalingModel.LinearTransformation(array)

		assert.Error(t, err2)
	})

	t.Run("Calling CreateModel multiple times should always return the same instance", func(t *testing.T) {
		testScalingModel, err := models.CreateModel("Scaling")
		assert.NoError(t, err)

		assert.Same(t, scalingModel, testScalingModel)
	})

	t.Run("Creating Reflection model should return correct transform matrix", func(t *testing.T) {
		reflectionModel, err := models.CreateModel("Reflection")
		assert.NoError(t, err)
		array := testArray(1.0, 1000)

		output, err2 := reflectionModel.LinearTransformation(array)
		assert.NoError(t, err2)

		for i := range output {
			assert.Equal(t, -1.0, output[i])
		}
		assert.Implements(t, (*mod.IModel)(nil), reflectionModel)
	})

	t.Run("Creating Shrinking model should return correct transform matrix", func(t *testing.T) {
		reflectionModel, err := models.CreateModel("Shrinking")
		assert.NoError(t, err)
		array := testArray(1.0, 1000)

		output, err2 := reflectionModel.LinearTransformation(array)
		assert.NoError(t, err2)

		for i := range output {
			assert.Equal(t, 0.5, output[i])
		}
		assert.Implements(t, (*mod.IModel)(nil), reflectionModel)
	})
}

// helpers
func testArray(value float64, size int) []float64 {
	//  creates an array of 1000 floats, all filled with the value 1.0
	arr := make([]float64, size)

	for i := 0; i < len(arr); i++ {
		arr[i] = value
	}

	return arr
}
