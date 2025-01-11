package framework

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestService struct{}

func NewTestService() *TestService {
	return &TestService{}
}

func TestContainer(t *testing.T) {
	container := NewContainer()

	container.Register("TestService", NewTestService(), Singleton)

	service := container.Get("TestService")
	assert.Equal(t, service, container.Get("TestService"))
}
