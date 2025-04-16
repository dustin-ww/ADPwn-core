package adpwn

type ModuleDependency struct {
	PreviousModule string `gorm:"column:previous_module" json:"previous_module"`
	NextModule     string `gorm:"column:next_module" json:"next_module"`
}
