package framework

import (
	"context"
	"fmt"
	"strings"
	"sync"
)

type ScopeStrategy interface {
	Init(def *ServiceDefinition)
	Resolve(c *Container, ctx context.Context, def *ServiceDefinition) any
	Cleanup(ctx context.Context)
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

func (s *SingletonStrategy) Cleanup(ctx context.Context) {
	// Singleton scope doesn't need cleanup
}

// PrototypeScope
type PrototypeStrategy struct{}

func (p *PrototypeStrategy) Init(def *ServiceDefinition) {
}

func (p *PrototypeStrategy) Resolve(_ *Container, _ context.Context, def *ServiceDefinition) any {
	return def.factory()
}

func (p *PrototypeStrategy) Cleanup(ctx context.Context) {
	// Prototype scope doesn't need cleanup
}

// HttpRequestScope
type HttpRequestScopeStrategy struct {
	instances sync.Map // key: requestID-serviceName, value: instance
}

func (h *HttpRequestScopeStrategy) Init(def *ServiceDefinition) {
}

func (h *HttpRequestScopeStrategy) Resolve(c *Container, ctx context.Context, def *ServiceDefinition) any {
	requestID := ctx.Value(REQUESTID)
	if requestID == nil {
		return def.factory()
	}

	instanceKey := fmt.Sprintf("%v-%s", requestID, def.name)
	if instance, ok := h.instances.Load(instanceKey); ok {
		return instance
	}

	instance := def.factory()
	h.instances.Store(instanceKey, instance)
	return instance
}

func (h *HttpRequestScopeStrategy) Cleanup(ctx context.Context) {
	requestID, ok := ctx.Value(REQUESTID).(string)
	if ok {
		h.clearRequestInstances(requestID)
	}
}

func (h *HttpRequestScopeStrategy) clearRequestInstances(requestID string) {
	prefix := fmt.Sprintf("%s-", requestID)
	h.instances.Range(func(key, value any) bool {
		if k, ok := key.(string); ok && strings.HasPrefix(k, prefix) {
			h.instances.Delete(k)
		}
		return true
	})
}
