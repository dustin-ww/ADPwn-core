package adpwn

import (
	"database/sql/driver"
	"encoding/json"
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

func (mt ModuleOptionType) Value() (driver.Value, error) {
	return mt.String(), nil
}

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

func (mt ModuleOptionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(mt.String())
}

func (mt *ModuleOptionType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	option, err := ParseModuleOptionType(s)
	if err != nil {
		return err
	}

	*mt = option
	return nil
}
