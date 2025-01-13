package framework

import (
	"context"
	"fmt"
	"sync"
	"strings"
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

func (h *HttpRequestScopeStrategy) ClearRequestInstances(requestID any) {
	h.instances.Range(func(key, value any) bool {
		if k, ok := key.(string); ok {
			if strings.HasPrefix(k, fmt.Sprintf("%v-", requestID)) {
				h.instances.Delete(key)
			}
		}
		return true
	})
}
