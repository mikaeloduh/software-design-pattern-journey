package framework

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestService struct {
	value int
}

func NewTestService() *TestService {
	return &TestService{value: 1}
}

func TestContainer_Singleton(t *testing.T) {
	t.Parallel()

	container := NewContainer()

	container.Register("TestService", func() any { return NewTestService() }, Singleton)

	testInstance := container.Get("TestService").(*TestService)
	expectedInstance := container.Get("TestService").(*TestService)

	assert.Same(t, expectedInstance, testInstance)
}

func TestContainer_Singleton_Parallel(t *testing.T) {
	container := NewContainer()
	container.Register("TestService", func() any { return NewTestService() }, Singleton)

	const concurrency = 100
	var wg sync.WaitGroup

	results := make([]*TestService, concurrency)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			results[i] = container.Get("TestService").(*TestService)
		}(i)
	}

	wg.Wait()

	for i := 0; i < concurrency; i++ {
		assert.Same(t, results[0], results[i])
	}
}

func TestContainer_Prototype(t *testing.T) {
	t.Parallel()

	container := NewContainer()

	container.Register("TestService", func() any { return NewTestService() }, Prototype)

	service1 := container.Get("TestService").(*TestService)
	service2 := container.Get("TestService").(*TestService)
	assert.NotSame(t, service1, service2)
}
