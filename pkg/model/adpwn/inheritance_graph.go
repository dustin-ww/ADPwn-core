package adpwn

type InheritanceGraph struct {
	Nodes []*Module           `json:"nodes"`
	Edges []*ModuleDependency `json:"edges"`
}
