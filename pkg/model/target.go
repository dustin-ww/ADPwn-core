package model

type Target struct {
	UID     string   `json:"uid,omitempty"`
	Name    string   `json:"name,omitempty"`
	IPRange string   `json:"ip_range,omitempty"`
	DType   []string `json:"dgraph.type,omitempty"`
}
