package model

import "time"

type Domain struct {
	UID                 string         `json:"uid,omitempty"`
	DNSName             string         `json:"dns_name,omitempty"`
	NetBiosName         string         `json:"net_bios_name,omitempty"`         // AD-NetBIOS-Name
	DomainGUID          string         `json:"domain_guid,omitempty"`           // domainGUID
	DomainSID           string         `json:"domain_sid,omitempty"`            // domainSID
	DomainFunctionLevel string         `json:"domain_function_level,omitempty"` // msDS-Behavior version
	ForestFunctionLevel string         `json:"forest_function_level,omitempty"`
	FSMORoleOwners      []string       `json:"fsmo_role_owners,omitempty"` // FSMO roles
	SecurityPolicies    SecurityPolicy `json:"security_policies,omitempty"`
	TrustRelationships  []Trust        `json:"trust_relationships,omitempty"`
	Created             time.Time      `json:"created,omitempty"`            // whenCreated
	LastModified        time.Time      `json:"last_modified,omitempty"`      // whenChanged
	LinkedGPOs          []string       `json:"linked_gpos,omitempty"`        // gPLink
	DefaultContainers   []string       `json:"default_containers,omitempty"` // wellKnownObjects
	BelongsToProject    Project        `json:"belongs_to_project,omitempty"`
	HasHost             []Host         `json:"has_host,omitempty"`
	HasUser             []User         `json:"has_user,omitempty"`
	DType               []string       `json:"dgraph.type,omitempty"`
}

type SecurityPolicy struct {
	MinPasswordLength int `json:"min_pwd_length,omitempty"`
	PasswordHistory   int `json:"pwd_history_length,omitempty"`
	LockoutThreshold  int `json:"lockout_threshold,omitempty"`
	LockoutDuration   int `json:"lockout_duration,omitempty"` // Minutes
}

type Trust struct {
	TrustedDomain string `json:"trusted_domain,omitempty"`
	Direction     string `json:"direction,omitempty"`  // inbound, outbound, bidirectional
	TrustType     string `json:"trust_type,omitempty"` // parent-child, cross-forest, external
	IsTransitive  bool   `json:"is_transitive,omitempty"`
}
