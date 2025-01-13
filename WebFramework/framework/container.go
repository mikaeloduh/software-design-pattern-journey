package framework

import (
	"sync"
)

type Factory func() any

type ServiceDefinition struct {
	name     string
	factory  Factory
	strategy ScopeStrategy
}

type Container struct {
	services map[string]*ServiceDefinition
	// key: *http.Request (or any other identifier you choose to use)
	// value: map[string]any (serviceName -> instance)
	scopeInstances sync.Map
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
	def, exists := c.services[name]
	if !exists {
		return nil
	}
	return def.strategy.Resolve(c, nil, def)
}

func (c *Container) GetWithContext(ctx ScopeContext, name string) any {
	def, exists := c.services[name]
	if !exists {
		return nil
	}
	return def.strategy.Resolve(c, ctx, def)
}

func (c *Container) ClearContext(ctx ScopeContext) {
	c.scopeInstances.Delete(ctx.ID())
}

// HttpRequestScopeMiddleware is a Middleware that clears the instance map for each request
func HttpRequestScopeMiddleware(container *Container) Middleware {
	return func(w *ResponseWriter, r *Request, next func()) error {
		// (1) entered Middleware, but not yet execute Handler
		//     at this point, if you need to do anything, you can do it here.
		//     for example, if you don't need it, just leave it blank.

		// (2) execute the next Middleware or final Handler
		next()

		// (3) after Handler execution, make sure to clear this request's instance map
		container.ClearContext(&HttpScopeContext{req: r.Request})

		// (4) return Handler execution result (if any)
		return nil
	}
}
