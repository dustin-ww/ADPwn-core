package adpwn

type ModuleOption struct {
	Key       string           `gorm:"column:option_key" json:"option_key"`
	Type      ModuleOptionType `gorm:"column:type" json:"option_type"`
	Required  bool             `gorm:"column:required" json:"option_required"`
	ModuleKey string           `gorm:"column:module_key" json:"module_key"`
}
