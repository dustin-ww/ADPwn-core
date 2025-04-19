// core/interfaces/module_executor.go
package interfaces

import (
	"ADPwn-core/pkg/model/adpwn/input"
)

type ModuleExecutor interface {
	ExecuteModule(key string, params *input.Parameter) error
}
