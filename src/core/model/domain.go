package model

import (
	"ADPwn/core/model/utils"
	"time"
)

type Domain struct {
	// Internal
	UID              string         `json:"uid,omitempty"`
	Name             string         `json:"name,omitempty"`
	BelongsToProject utils.UIDRef   `json:"belongs_to_project,omitempty"`
	HasHost          []utils.UIDRef `json:"has_host,omitempty"`
	HasUser          []utils.UIDRef `json:"has_user,omitempty"`
	DType            []string       `json:"dgraph.type,omitempty"`
	// AD related
	DNSName             string         `json:"dns_name,omitempty"`
	NetBiosName         string         `json:"net_bios_name,omitempty"`
	DomainGUID          string         `json:"domain_guid,omitempty"`
	DomainSID           string         `json:"domain_sid,omitempty"`
	DomainFunctionLevel string         `json:"domain_function_level,omitempty"`
	ForestFunctionLevel string         `json:"forest_function_level,omitempty"`
	FSMORoleOwners      []string       `json:"fsmo_role_owners,omitempty"`
	SecurityPolicies    utils.UIDRef   `json:"security_policies,omitempty"`
	TrustRelationships  []utils.UIDRef `json:"trust_relationships,omitempty"`
	Created             time.Time      `json:"created,omitempty"`
	LastModified        time.Time      `json:"last_modified,omitempty"`
	LinkedGPOs          []string       `json:"linked_gpos,omitempty"`
	DefaultContainers   []string       `json:"default_containers,omitempty"`
}

type SecurityPolicy struct {
	MinPasswordLength int `json:"min_pwd_length,omitempty"`
	PasswordHistory   int `json:"pwd_history_length,omitempty"`
	LockoutThreshold  int `json:"lockout_threshold,omitempty"`
	LockoutDuration   int `json:"lockout_duration,omitempty"`
}

type Trust struct {
	TrustedDomain string `json:"trusted_domain,omitempty"`
	Direction     string `json:"direction,omitempty"`  // inbound, outbound, bidirectional
	TrustType     string `json:"trust_type,omitempty"` // parent-child, cross-forest, external
	IsTransitive  bool   `json:"is_transitive,omitempty"`
}
