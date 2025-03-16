package model

type ADPwnModule struct {
	UID         string   `json:"uid,omitempty"`
	AttackID    string   `json:"attack_id,omitempty"`
	Metric      string   `json:"metric,omitempty"`
	Description string   `json:"description,omitempty"`
	Name        string   `json:"name,omitempty"`
	Version     string   `json:"version,omitempty"`
	Author      string   `json:"author,omitempty"`
	DType       []string `json:"dgraph.type,omitempty"`
	IsAttack    bool     `json:"is_attack,omitempty"`
}

type ADPwnModuleMetadata struct {
	UID           string        `json:"uid,omitempty"`
	AttackID      string        `json:"attack_id,omitempty"`
	Metric        string        `json:"metric,omitempty"`
	Name          string        `json:"name,omitempty"`
	Version       string        `json:"version,omitempty"`
	Author        string        `json:"Author,omitempty"`
	LastRun       string        `json:"last_run,omitempty"`
	HasDependency []ADPwnModule `json:"has_dependency,omitempty"`
	IsAssumedBy   []ADPwnModule `json:"is_assumed_by,omitempty"`
	DType         []string      `json:"dgraph.type,omitempty"`
}
