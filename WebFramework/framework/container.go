package framework

type Scope int

const (
	Singleton Scope = iota
	Prototype
)

type Factory func() any

type ServiceDefinition struct {
	factory Factory
	scope   Scope
}

type Container struct {
	services map[string]*ServiceDefinition
}

func NewContainer() *Container {
	return &Container{
		services: make(map[string]*ServiceDefinition),
	}
}

func (c *Container) Register(name string, factory Factory, scope Scope) {
	if scope == Singleton {
		instance := factory()
		c.services[name] = &ServiceDefinition{
			factory: func() any {
				return instance
			},
			scope: scope,
		}
	} else {
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
