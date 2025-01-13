package framework

import "net/http"

type ScopeStrategy interface {
	Init(def *ServiceDefinition)
	Resolve(c *Container, r *http.Request, def *ServiceDefinition) any
}

// SingletonScope
type SingletonStrategy struct {
	instance any
}

func (s *SingletonStrategy) Init(def *ServiceDefinition) {
	s.instance = def.factory()
}

func (s *SingletonStrategy) Resolve(_ *Container, _ *http.Request, _ *ServiceDefinition) any {
	return s.instance
}

// PrototypeScope
type PrototypeStrategy struct{}

func (p *PrototypeStrategy) Init(def *ServiceDefinition) {
}

func (p *PrototypeStrategy) Resolve(_ *Container, _ *http.Request, def *ServiceDefinition) any {
	return def.factory()
}

// HttpRequestScope
type HttpRequestScopeStrategy struct{}

func (h *HttpRequestScopeStrategy) Init(def *ServiceDefinition) {
}

func (h *HttpRequestScopeStrategy) Resolve(c *Container, r *http.Request, def *ServiceDefinition) any {
	if r == nil {
		panic("request is nil")
	}

	val, _ := c.requestInstances.LoadOrStore(r, make(map[string]any))
	instanceMap := val.(map[string]any)

	if inst, ok := instanceMap[def.name]; ok {
		return inst
	}
	newInst := def.factory()
	instanceMap[def.name] = newInst
	return newInst
}
