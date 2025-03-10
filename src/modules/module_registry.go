package modules

type Registry struct {
	modules []ADPwnModule
}

// NewRegistry creates a new module registry
func NewRegistry() *Registry {
	return &Registry{
		modules: make([]ADPwnModule, 0),
	}
}

// RegisterEnumerationModule adds a new enumeration module to the registry
func (r *Registry) RegisterEnumerationModule(module ADPwnModule) {
	r.modules = append(r.modules, module)
}

// GetModules returns all registered modules
func (r *Registry) GetModules() []ADPwnModule {
	return r.modules
}

// Global registry instance
var GlobalRegistry = NewRegistry()
