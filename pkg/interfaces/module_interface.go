package interfaces

import (
	"ADPwn-core/internal/sse"
	"ADPwn-core/pkg/model/adpwn"
)

type ADPwnModule interface {
	GetConfigKey() string
	ExecuteModule(params *adpwn.Parameter, logger *sse.SSELogger) error
}
