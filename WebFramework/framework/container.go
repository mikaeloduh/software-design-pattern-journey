package framework

type Scope int

const (
	Singleton Scope = iota
	Transient
)

type ServiceDefinition struct {
	instance interface{}
	scope    Scope
}

type Container struct {
	services map[string]*ServiceDefinition
}

func NewContainer() *Container {
	return &Container{
		services: make(map[string]*ServiceDefinition),
	}
}

func (c *Container) Register(name string, service interface{}, scope Scope) {
	c.services[name] = &ServiceDefinition{
		instance: service,
		scope:    scope,
	}
}

func (c *Container) Get(name string) interface{} {
	if service, exists := c.services[name]; exists {
		return service.instance
	}
	return nil
}