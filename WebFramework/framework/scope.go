package framework

import (
	"context"
	"fmt"
)

type ScopeStrategy interface {
	Init(def *ServiceDefinition)
	Resolve(c *Container, ctx context.Context, def *ServiceDefinition) any
}

// SingletonScope
type SingletonStrategy struct {
	instance any
}

func (s *SingletonStrategy) Init(def *ServiceDefinition) {
	s.instance = def.factory()
}

func (s *SingletonStrategy) Resolve(_ *Container, _ context.Context, _ *ServiceDefinition) any {
	return s.instance
}

// PrototypeScope
type PrototypeStrategy struct{}

func (p *PrototypeStrategy) Init(def *ServiceDefinition) {
}

func (p *PrototypeStrategy) Resolve(_ *Container, _ context.Context, def *ServiceDefinition) any {
	return def.factory()
}

// HttpRequestScope
type HttpRequestScopeStrategy struct{}

func (h *HttpRequestScopeStrategy) Init(def *ServiceDefinition) {
}

func (h *HttpRequestScopeStrategy) Resolve(c *Container, ctx context.Context, def *ServiceDefinition) any {
	requestID := ctx.Value(REQUESTID)
	if requestID == nil {
		return def.factory()
	}

	instanceKey := fmt.Sprintf("%v-%s", requestID, def.name)
	if instance, ok := c.scopeInstances.Load(instanceKey); ok {
		return instance
	}

	instance := def.factory()
	c.scopeInstances.Store(instanceKey, instance)
	return instance
}
