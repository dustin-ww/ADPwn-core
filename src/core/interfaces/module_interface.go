package interfaces

import (
	"ADPwn/core/model/adpwn"
)

type ADPwnModule interface {
	GetConfigKey() string
	Execute(project adpwn.Parameter) error
}
