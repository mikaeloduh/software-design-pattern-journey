package framework

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
)

type Factory func() any

type ServiceDefinition struct {
	name     string
	factory  Factory
	strategy ScopeStrategy
}

type Container struct {
	services       map[string]*ServiceDefinition
	scopeInstances sync.Map // key: InstaceKey, value: instance
}

func NewContainer() *Container {
	return &Container{
		services: make(map[string]*ServiceDefinition),
	}
}

func (c *Container) Register(name string, factory Factory, strategy ScopeStrategy) {
	def := &ServiceDefinition{
		name:     name,
		factory:  factory,
		strategy: strategy,
	}

	def.strategy.Init(def)
	c.services[name] = def
}

func (c *Container) Get(name string) any {
	return c.GetWithContext(context.Background(), name)
}

func (c *Container) GetWithContext(ctx context.Context, name string) any {
	def, exists := c.services[name]
	if !exists {
		return nil
	}
	return def.strategy.Resolve(c, ctx, def)
}

type InstaceKey string

const REQUESTID = InstaceKey("request_id")

// HttpRequestScopeMiddleware is a Middleware that manages request scoped services
func HttpRequestScopeMiddleware(container *Container) Middleware {
	var requestCounter uint64

	return func(w *ResponseWriter, r *Request, next func()) error {
		requestID := atomic.AddUint64(&requestCounter, 1)

		ctx := context.WithValue(r.Context(), REQUESTID, requestID)
		r.Request = r.Request.WithContext(ctx)

		next()

		// clean up all instances associated with this request
		container.scopeInstances.Range(func(key, value any) bool {
			if k, ok := key.(string); ok {
				if strings.HasPrefix(k, fmt.Sprintf("%v-", requestID)) {
					container.scopeInstances.Delete(key)
				}
			}
			return true
		})

		return nil
	}
}
