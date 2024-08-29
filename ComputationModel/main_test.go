package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Main(t *testing.T) {
	models := NewModels()

	assert.IsType(t, &Models{}, models)

	t.Run("test scaling model", func(t *testing.T) {
		scalingModel, err := models.CreateModel("Scaling")
		assert.NoError(t, err)

		input := testArray()

		output := scalingModel.LinearTransformation(input)

		assert.Equal(t, 2.0, output[0])
		assert.Equal(t, 2.0, output[1])
		assert.Implements(t, (*IModel)(nil), scalingModel)
	})
}

// helpers
func testArray() []float64 {
	//  creates an array of 1000 floats, all filled with the value 1.0
	arr := make([]float64, 1000)

	for i := 0; i < len(arr); i++ {
		arr[i] = 1.0
	}

	return arr
}
