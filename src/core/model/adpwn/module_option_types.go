package adpwn

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

type ModuleOptionType int

const (
	Checkbox ModuleOptionType = iota
	TextInput
	UserSelection
	TargetSelection
)

var moduleOptionTypeMap = map[string]ModuleOptionType{
	"checkbox":        Checkbox,
	"textInput":       TextInput,
	"userSelection":   UserSelection,
	"targetSelection": TargetSelection,
}

func (mt ModuleOptionType) String() string {
	names := [...]string{"checkbox", "textInput", "userSelection", "targetSelection"}
	if int(mt) < len(names) {
		return names[mt]
	}
	return "Unknown Option Type"
}

func ParseModuleOptionType(moduleOptionStr string) (ModuleOptionType, error) {
	if optionType, exists := moduleOptionTypeMap[moduleOptionStr]; exists {
		return optionType, nil
	}
	return 0, errors.New("invalid module option type: " + moduleOptionStr)
}

// Value implementiert die driver.Valuer-Schnittstelle, sodass der Enum als String in die DB geschrieben werden kann.
func (mt ModuleOptionType) Value() (driver.Value, error) {
	return mt.String(), nil
}

// Scan implementiert die sql.Scanner-Schnittstelle, sodass der DB-Wert korrekt in den Enum umgewandelt wird.
func (mt *ModuleOptionType) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("cannot scan %v into ModuleOptionType", value)
	}
	option, err := ParseModuleOptionType(str)
	if err != nil {
		return err
	}
	*mt = option
	return nil
}
