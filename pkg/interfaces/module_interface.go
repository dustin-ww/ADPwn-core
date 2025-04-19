package interfaces

import (
	"ADPwn-core/internal/sse"
	"ADPwn-core/pkg/model/adpwn/input"
)

type ADPwnModule interface {
	GetConfigKey() string
	ExecuteModule(params *input.Parameter, logger *sse.SSELogger) error
}
