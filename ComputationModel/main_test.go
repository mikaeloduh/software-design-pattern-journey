package main

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"computationmodel/model"
)

var modelsInstance model.IModels
var scalingModel model.IModel

func Test_Main(t *testing.T) {
	t.Parallel()

	modelsInstance = model.NewModels()
	models := model.NewModels()

	assert.Same(t, modelsInstance, models)

	t.Run("Creating Scaling model should return correct transform matrix", func(t *testing.T) {
		var err error
		scalingModel, err = models.CreateModel("Scaling")
		assert.NoError(t, err)

		array := testArray(1.0, 1000)

		output, err2 := scalingModel.LinearTransformation(array)
		assert.NoError(t, err2)

		for i := range output {
			assert.Equal(t, 2.0, output[i])
		}
		assert.Implements(t, (*model.IModel)(nil), scalingModel)
	})

	t.Run("Validating array length must equal the model's row size", func(t *testing.T) {
		var err error
		scalingModel, err = models.CreateModel("Scaling")
		assert.NoError(t, err)
		array := testArray(1.0, 999)

		_, err2 := scalingModel.LinearTransformation(array)

		assert.Error(t, err2)
	})

	t.Run("Calling CreateModel multiple times should always return the same instance", func(t *testing.T) {
		actualScalingModel, err := models.CreateModel("Scaling")
		assert.NoError(t, err)

		assert.Same(t, scalingModel, actualScalingModel)
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
		assert.Implements(t, (*model.IModel)(nil), reflectionModel)
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
		assert.Implements(t, (*model.IModel)(nil), reflectionModel)
	})
}

func Test_Async(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Create a new subtest for each goroutine to ensure proper test isolation
			t.Run(fmt.Sprintf("async-%02d", i), func(t *testing.T) {
				Test_Main(t)
			})
		}(i)
	}

	wg.Wait()
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
