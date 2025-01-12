package framework

import (
	"net/http"
	"sync"
)

type Scope int

const (
	SingletonScope Scope = iota
	PrototypeScope
	HttpRequestScope
)

type Factory func() any

type ServiceDefinition struct {
	factory Factory
	scope   Scope
}

type Container struct {
	services map[string]*ServiceDefinition
	// key: *http.Request (or any other identifier you choose to use)
	// value: map[string]any (serviceName -> instance)
	requestInstances sync.Map
}

func NewContainer() *Container {
	return &Container{
		services: make(map[string]*ServiceDefinition),
	}
}

func (c *Container) Register(name string, factory Factory, scope Scope) {
	switch scope {
	case SingletonScope:
		instance := factory()
		c.services[name] = &ServiceDefinition{
			factory: func() any { return instance },
			scope:   scope,
		}
	case HttpRequestScope:
		c.services[name] = &ServiceDefinition{
			factory: factory,
			scope:   scope,
		}
	default: // PrototypeScope
		c.services[name] = &ServiceDefinition{
			factory: factory,
			scope:   scope,
		}
	}
}

func (c *Container) Get(name string) any {
	if service, exists := c.services[name]; exists {
		return service.factory()
	}
	return nil
}

// for HttpRequestScope, we need to be able to receive requests
func (c *Container) GetFromRequest(r *http.Request, name string) any {
	def, ok := c.services[name]
	if !ok {
		return nil
	}

	if def.scope != HttpRequestScope {
		return def.factory()
	}

	// HttpRequestScope: lookup instance map
	val, _ := c.requestInstances.LoadOrStore(r, make(map[string]any))
	instanceMap := val.(map[string]any)

	if inst, found := instanceMap[name]; found {
		// service already exists, return it
		return inst
	}

	// instance doesn't exist, create it
	newInst := def.factory()
	instanceMap[name] = newInst
	return newInst
}

func (c *Container) ClearRequest(r *http.Request) {
	c.requestInstances.Delete(r)
}

func HttpRequestScopeMiddleware(container *Container) Middleware {
	return func(w *ResponseWriter, r *Request, next func()) error {
		// (1) entered Middleware, but not yet execute Handler
		//     at this point, if you need to do anything, you can do it here.
		//     for example, if you don't need it, just leave it blank.

		// (2) execute the next Middleware or final Handler
		next()

		// (3) after Handler execution, make sure to clear this request's instance map
		container.ClearRequest(r.Request)

		// (4) return Handler execution result (if any)
		return nil
	}
}
