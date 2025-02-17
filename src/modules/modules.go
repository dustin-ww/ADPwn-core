package modules

import (
	_ "ADPwn/modules/enumeration"
	"ADPwn/modules/internal/base"
)

func GetADPwnModules() []base.ADPwnModule {
	return base.GlobalRegistry.GetModules()
}
