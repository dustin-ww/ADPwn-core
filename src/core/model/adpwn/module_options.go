package adpwn

type ModuleOptionType int

const (
	Checkbox ModuleOptionType = iota
	Textfield
	UserSelection
	TargetSelection
)

// String mit Pointer-Empfänger
func (mt *ModuleOptionType) String() string {
	if mt == nil {
		return "UnknownModule"
	}
	return [...]string{"EnumerationModule", "AttackModule"}[*mt]
}
