package framework

import (
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

	container.Register("TestService", func() any {
		return NewTestService()
	}, Singleton)

	service := container.Get("TestService").(*TestService)
	assert.Same(t, service, container.Get("TestService"))
}

func TestContainer_Prototype(t *testing.T) {
	t.Parallel()

	container := NewContainer()

	container.Register("TestService", func() any {
		return NewTestService()
	}, Prototype)

	service1 := container.Get("TestService").(*TestService)
	service2 := container.Get("TestService").(*TestService)
	assert.NotSame(t, service1, service2)
}
