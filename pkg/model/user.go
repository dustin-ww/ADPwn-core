package model

type User struct {
	UID             string   `json:"uid,omitempty"`
	Name            string   `json:"name,omitempty"`
	NTLMHash        string   `json:"ntlm_hash,omitempty"`
	Password        string   `json:"password,omitempty"`
	IsAdmin         bool     `json:"is_admin,omitempty"`
	BelongsToDomain Domain   `json:"belongs_to_domain,omitempty"`
	DType           []string `json:"dgraph.type,omitempty"`
}
