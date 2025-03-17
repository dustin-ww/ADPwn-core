package adpwn

import "ADPwn/core/interfaces"

type Module struct {
	AttackID        string                `gorm:"column:attack_id" json:"attack_id"`
	ExecutionMetric string                `gorm:"column:execution_metric" json:"execution_metric"`
	Description     string                `gorm:"column:description" json:"description"`
	Name            string                `gorm:"column:name" json:"name"`
	Version         string                `gorm:"column:version" json:"version"`
	Author          string                `gorm:"column:author" json:"author"`
	ModuleType      interfaces.ModuleType `gorm:"column:module_type" json:"module_type"`
	LootPath        string                `gorm:"column:loot_path" json:"loot_path"`
	Key             string                `gorm:"column:key" json:"key"`
}

type ModuleInheritanceEdge struct {
	PreviousModule string `gorm:"column:previous_module" json:"previous_module"`
	NextModule     string `gorm:"column:next_module" json:"next_module"`
}
