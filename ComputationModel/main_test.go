package main

import (
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

func TestModels_Concurrency(t *testing.T) {
	const concurrency = 100
	var wg sync.WaitGroup

	instances := make([]model.IModels, concurrency)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			instances[i] = model.NewModels()
		}(i)
	}

	wg.Wait()

	for i := 0; i < concurrency; i++ {
		assert.Same(t, instances[0], instances[i])
	}
}

func TestModel_Concurrency(t *testing.T) {
	const concurrency = 100
	var wg sync.WaitGroup

	models := model.NewModels()

	instances := make([]model.IModel, concurrency+1)

	instances[0], _ = models.CreateModel("Scaling")

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			var err error
			instances[i+1], err = models.CreateModel("Scaling")
			assert.NoError(t, err)
		}(i)
	}

	wg.Wait()

	for i := 0; i < concurrency+1; i++ {
		assert.Same(t, instances[0], instances[i])
	}
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
