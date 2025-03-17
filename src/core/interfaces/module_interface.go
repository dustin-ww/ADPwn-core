package interfaces

import (
	"database/sql/driver"
	"fmt"
)

type ModuleType int

const (
	EnumerationModule ModuleType = iota
	AttackModule
)

func (mt ModuleType) String() string {
	return [...]string{"EnumerationModule", "AttackModule"}[mt]
}

func ParseModuleType(s string) (ModuleType, error) {
	switch s {
	case "AttackModule":
		return AttackModule, nil
	case "EnumerationModule":
		return EnumerationModule, nil
	default:
		return 0, fmt.Errorf("invalid ModuleType: %s", s)
	}
}

func (m *ModuleType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("erwartete string für ModuleType, bekam: %T", value)
	}

	parsed, err := ParseModuleType(strValue)
	if err != nil {
		return err
	}
	*m = parsed
	return nil
}

func (m ModuleType) Value() (driver.Value, error) {
	return m.String(), nil
}

type ADPwnModule interface {
	GetConfigKey() string
}
