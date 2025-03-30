package repository

import (
	"ADPwn/core/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

type DomainRepository interface {
	//CRUD
	Create(ctx context.Context, name string) (string, error) // Returns UID
	Get(ctx context.Context, tx *dgo.Txn, uid string) (*model.Domain, error)
	UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error
	CreateWithObject(ctx context.Context, tx *dgo.Txn, model *model.Domain) (string, error)
	//Relations
	AddHost(ctx context.Context, domainUID, hostUID string) error
	AddUser(ctx context.Context, domainUID, userUID string) error
	GetByProjectUID(ctx context.Context, tx *dgo.Txn, projectUID string) ([]*model.Domain, error)
}

type DgraphDomainRepository struct {
	DB *dgo.Dgraph
}

func (r *DgraphDomainRepository) CreateWithObject(ctx context.Context, tx *dgo.Txn, domain *model.Domain) (string, error) {
	domain.DType = []string{"Domain"}

	jsonData, err := json.Marshal(domain)
	if err != nil {
		return "", fmt.Errorf("marshal error: %w", err)
	}

	mu := &api.Mutation{
		SetJson: jsonData,
	}

	assigned, err := tx.Mutate(ctx, mu)
	if err != nil {
		return "", fmt.Errorf("mutation error: %w", err)
	}

	return assigned.Uids["blank-0"], nil
}

func (r *DgraphDomainRepository) Get(ctx context.Context, tx *dgo.Txn, uid string) (*model.Domain, error) {
	query := `
        query Domain($uid: string) {
            domain(func: uid($uid)) {
                uid
                name
                belongs_to_project
                has_hosts {
                    uid
				}
				has_user {
					uid
				}

            }
        }`
	vars := map[string]string{"$uid": uid}
	res, err := tx.QueryWithVars(ctx, query, vars)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	var result struct {
		Domain []model.Domain `json:"domain"`
	}
	if err := json.Unmarshal(res.Json, &result); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	if len(result.Domain) == 0 {
		return nil, fmt.Errorf("domain not found: %s", uid)
	}

	return &result.Domain[0], nil
}

func (r *DgraphDomainRepository) GetByProjectUID(ctx context.Context, tx *dgo.Txn, projectUID string) ([]*model.Domain, error) {
	query := `
        query DomainsByProject($projectUID: string) {
            domains(func: has(belongs_to_project)) @filter(uid_in(belongs_to_project, $projectUID)) {
                uid
                dns_name
                net_bios_name
                domain_guid
                domain_sid
                domain_function_level
                forest_function_level
                fsmo_role_owners
                created
                last_modified
                linked_gpos
                default_containers
                dgraph.type
                
                security_policies {
                    min_pwd_length
                    pwd_history_length
                    lockout_threshold
                    lockout_duration
                }
                
                trust_relationships {
                    trusted_domain
                    direction
                    trust_type
                    is_transitive
                }

                belongs_to_project {
                    uid
                }

                has_host {
                    uid
                    
                }
                
                has_user {
                    uid
                }
            }
        }
    `

	vars := map[string]string{"$projectUID": projectUID}
	res, err := tx.QueryWithVars(ctx, query, vars)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	var result struct {
		Domains []*model.Domain `json:"domains"`
	}
	if err := json.Unmarshal(res.Json, &result); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	//for _, domain := range result.Domains {
	//	// domain.Created = formatTime(domain.RawCreated)
	//}

	return result.Domains, nil
}

func (r *DgraphDomainRepository) Create(ctx context.Context, name string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphDomainRepository) UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphDomainRepository) AddHost(ctx context.Context, domainUID, hostUID string) error {
	//TODO implement me
	panic("implement me")
}

func (r *DgraphDomainRepository) AddUser(ctx context.Context, domainUID, userUID string) error {
	//TODO implement me
	panic("implement me")
}

func NewDgraphDomainRepository(db *dgo.Dgraph) *DgraphDomainRepository {
	return &DgraphDomainRepository{DB: db}
}
