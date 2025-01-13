package framework

type ScopeStrategy interface {
	Init(def *ServiceDefinition)
	Resolve(c *Container, ctx ScopeContext, def *ServiceDefinition) any
}

// SingletonScope
type SingletonStrategy struct {
	instance any
}

func (s *SingletonStrategy) Init(def *ServiceDefinition) {
	s.instance = def.factory()
}

func (s *SingletonStrategy) Resolve(_ *Container, _ ScopeContext, _ *ServiceDefinition) any {
	return s.instance
}

// PrototypeScope
type PrototypeStrategy struct{}

func (p *PrototypeStrategy) Init(def *ServiceDefinition) {
}

func (p *PrototypeStrategy) Resolve(_ *Container, _ ScopeContext, def *ServiceDefinition) any {
	return def.factory()
}

// HttpRequestScope
type HttpRequestScopeStrategy struct{}

func (h *HttpRequestScopeStrategy) Init(def *ServiceDefinition) {
	// do nothing
}

func (h *HttpRequestScopeStrategy) Resolve(c *Container, ctx ScopeContext, def *ServiceDefinition) any {
	// ctx should be HttpScopeContext
	// if ctx is not HttpScopeContext, you should handle it (panic or default to Prototype)
	httpCtx, ok := ctx.(*HttpScopeContext)
	if !ok {
		return def.factory()
	}

	// container may need to be changed to c.scopeInstances
	// here we use httpCtx.ID() as key
	val, _ := c.scopeInstances.LoadOrStore(httpCtx.ID(), make(map[string]any))
	instanceMap := val.(map[string]any)

	if inst, found := instanceMap[def.name]; found {
		return inst
	}

	newInst := def.factory()
	instanceMap[def.name] = newInst
	return newInst
}
