package model

import (
	"time"
)

type Project struct {
	// Internal
	UID         string    `json:"uid,omitempty"`
	Name        string    `json:"name,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
	Description string    `json:"description,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	UpdatedAt   time.Time `json:"updatedAt,omitempty"`
	HasTarget   []Target  `json:"has_target,omitempty"`
	HasDomain   []Domain  `json:"has_domain,omitempty"`
	DType       []string  `json:"dgraph.type,omitempty"`
}
