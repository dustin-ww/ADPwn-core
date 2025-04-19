package plugin

import (
	"ADPwn-core/internal/config"
	"ADPwn-core/internal/sse"
	"ADPwn-core/pkg/interfaces"
	"ADPwn-core/pkg/model/adpwn"
	"ADPwn-core/pkg/model/adpwn/input"
	"ADPwn-core/pkg/service"
	"context"
	"fmt"
	"log"
)

type Registry struct {
	modules         map[string]*adpwn.Module
	implementations map[string]interfaces.ADPwnModule
}

var GlobalRegistry = &Registry{
	modules:         make(map[string]*adpwn.Module),
	implementations: make(map[string]interfaces.ADPwnModule),
}

func RegisterPlugin(module interfaces.ADPwnModule) {
	log.Printf("--- Starting the ADPwn Module Loading Process...")
	configKey := module.GetConfigKey()
	modules, inherits, err := config.ModuleFromConfig(configKey)
	if err != nil {
		panic("register plugin fail:" + err.Error())
	}

	GlobalRegistry.modules[modules.Key] = modules
	GlobalRegistry.implementations[modules.Key] = module

	log.Printf("--- Finished the ADPwn Module Loading Process. Registering modulelib in Database...")
	handoverToService(modules, inherits)
}

func GetAll() []*adpwn.Module {
	modules := make([]*adpwn.Module, 0, len(GlobalRegistry.modules))
	for _, module := range GlobalRegistry.modules {
		modules = append(modules, module)
	}
	return modules
}

func GetModule(key string) *adpwn.Module {
	return GlobalRegistry.modules[key]
}

// Make sure Registry implements interfaces.ModuleExecutor
// registry.go
func (r *Registry) ExecuteModule(key string, params *input.Parameter) error {
	impl, ok := r.implementations[key]
	if !ok {
		return fmt.Errorf("no implementation found for module key: %s", key)
	}

	// Haupt-Logger mit RunID erstellen
	baseLogger := sse.GetLogger(params.RunID)

	// Modulspezifischen Logger erstellen
	moduleLogger := baseLogger.ForModule(impl.GetConfigKey())

	return impl.ExecuteModule(params, moduleLogger)
}

func ExecuteModule(key string, params *input.Parameter) error {
	return GlobalRegistry.ExecuteModule(key, params)
}

func handoverToService(module *adpwn.Module, inherits []*adpwn.ModuleDependency) {
	moduleService, err := service.NewADPwnModuleService(nil)
	if err != nil {
		err = fmt.Errorf("failed to create project service: %v", err)
		fmt.Println(err)
		return
	}

	_, err = moduleService.CreateWithObject(context.Background(), module)
	if err != nil {
		fmt.Println("failed to register plugin in db: " + err.Error())
	}

	err = moduleService.CreateModuleInheritanceEdges(context.Background(), inherits)
	if err != nil {
		fmt.Println("failed to register inheritance reference in db: " + err.Error())
	}
}
