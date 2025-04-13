package adpwn

type ModuleOption struct {
	Key         string           `gorm:"column:option_key" json:"key"`
	Type        ModuleOptionType `gorm:"column:type" json:"type"`
	Required    bool             `gorm:"column:required" json:"required"`
	ModuleKey   string           `gorm:"column:module_key" json:"-"`
	Label       string           `gorm:"column:label" json:"label"`
	Placeholder string           `gorm:"column:placeholder" json:"placeholder"`
}
