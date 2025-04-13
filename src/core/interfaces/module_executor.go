// core/interfaces/module_executor.go
package interfaces

import "ADPwn/core/model/adpwn"

type ModuleExecutor interface {
	ExecuteModule(key string, params *adpwn.Parameter) error
}
