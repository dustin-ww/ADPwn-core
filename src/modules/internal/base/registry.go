package base

var (
	GlobalRegistry = NewModuleRegistry()
)

type ModuleRegistry struct {
	modules             []ADPwnModule
	enumerationModules  []ADPwnModule
	exploitationModules []ADPwnModule
}

func NewModuleRegistry() *ModuleRegistry {
	return &ModuleRegistry{
		modules: make([]ADPwnModule, 0),
	}
}

func (r *ModuleRegistry) RegisterEnumerationModule(module ADPwnModule) {
	r.modules = append(r.modules, module)
	r.enumerationModules = append(r.enumerationModules, module)
}

func (r *ModuleRegistry) RegisterExploitationModule(module ADPwnModule) {
	r.modules = append(r.modules, module)
	r.exploitationModules = append(r.exploitationModules, module)
}

func (r *ModuleRegistry) GetModules() []ADPwnModule {
	return r.modules
}

func (r *ModuleRegistry) GetEnumerationModules() []ADPwnModule {
	return r.enumerationModules
}

func (r *ModuleRegistry) GetExploitationModules() []ADPwnModule {
	return r.exploitationModules
}
