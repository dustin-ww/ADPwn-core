package adpwn

type Collection struct {
	ID          uint     `gorm:"column:id" json:"id"`
	Name        string   `gorm:"column:name" json:"name"`
	Description string   `gorm:"column:description" json:"description"`
	Modules     []Module `json:"modules"`
}
