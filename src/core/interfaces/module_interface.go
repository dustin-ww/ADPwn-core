package interfaces

import (
	"ADPwn/core/model/adpwn"
	"ADPwn/sse/sse"
)

type ADPwnModule interface {
	GetConfigKey() string
	ExecuteModule(params *adpwn.Parameter, logger *sse.SSELogger) error
}
