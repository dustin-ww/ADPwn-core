// core/interfaces/module_executor.go
package interfaces

import "ADPwn-core/pkg/model/adpwn"

type ModuleExecutor interface {
	ExecuteModule(key string, params *adpwn.Parameter) error
}
