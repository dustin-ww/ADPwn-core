package plugin

import (
	"ADPwn/adapter/config"
	"ADPwn/core/interfaces"
	"ADPwn/core/model/adpwn"
	"ADPwn/core/service"
	"context"
	"fmt"
)

type Registry struct {
	modules []*adpwn.Module
}

var GlobalRegistry = &Registry{
	modules: make([]*adpwn.Module, 0),
}

func RegisterPlugin(module interfaces.ADPwnModule) {
	modules, inherits, err := config.ModuleFromConfig(module.GetConfigKey())
	if err != nil {
		panic("register plugin fail:" + err.Error())
	}
	GlobalRegistry.modules = append(GlobalRegistry.modules, modules)
	handoverToService(modules, inherits)
}

func GetAll() []*adpwn.Module {
	return GlobalRegistry.modules
}

func handoverToService(module *adpwn.Module, inherits []*adpwn.ModuleDependency) {
	moduleService, err := service.NewADPwnModuleService()
	if err != nil {
		err = fmt.Errorf("failed to create project service: %v", err)
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
