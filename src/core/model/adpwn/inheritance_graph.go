package adpwn

type InheritanceGraph struct {
	Nodes []*Module                `json:"nodes"`
	Edges []*ModuleInheritanceEdge `json:"edges"`
}
