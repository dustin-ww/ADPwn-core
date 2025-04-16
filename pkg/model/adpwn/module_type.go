package adpwn

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type ModuleType int

const (
	EnumerationModule ModuleType = iota
	AttackModule
)

// String mit Pointer-Empfänger
func (mt ModuleType) String() string {
	return [...]string{"EnumerationModule", "AttackModule"}[mt]
}

// Parse-Funktion bleibt gleich
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

// Scan weiterhin mit Pointer
func (mt *ModuleType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	strValue, ok := value.(string)
	if !ok {
		return fmt.Errorf("expected string for ModuleType, got: %T", value)
	}

	parsed, err := ParseModuleType(strValue)
	if err != nil {
		return err
	}
	*mt = parsed
	return nil
}

// Value mit Pointer-Empfänger
func (mt ModuleType) Value() (driver.Value, error) {
	return mt.String(), nil
}

// MarshalJSON mit Pointer-Empfänger
func (mt *ModuleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(mt.String())
}

// UnmarshalJSON bleibt gleich
func (mt *ModuleType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed, err := ParseModuleType(s)
	if err != nil {
		return err
	}
	*mt = parsed
	return nil
}
