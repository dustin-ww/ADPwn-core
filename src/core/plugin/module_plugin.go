package plugin

import "ADPwn/core/interfaces"

type Registry struct {
	enumerationModules []interfaces.ADPwnModule
	attackModules      []interfaces.ADPwnModule
}

var GlobalRegistry = &Registry{
	enumerationModules: make([]interfaces.ADPwnModule, 0),
	attackModules:      make([]interfaces.ADPwnModule, 0),
}

func RegisterEnumeration(module interfaces.ADPwnModule) {
	GlobalRegistry.enumerationModules = append(GlobalRegistry.enumerationModules, module)
}

func RegisterAttack(module interfaces.ADPwnModule) {
	GlobalRegistry.attackModules = append(GlobalRegistry.attackModules, module)
}

// GetAll gibt alle registrierten Module zurück
func GetAll() []interfaces.ADPwnModule {
	return append(GlobalRegistry.attackModules, GlobalRegistry.enumerationModules...)
}
