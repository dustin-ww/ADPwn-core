package repository

import (
	"ADPwn/core/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

type ProjectRepository interface {
	//CRUD
	Create(ctx context.Context, name string) (string, error) // Returns UID
	Get(ctx context.Context, uid string) (*model.Project, error)
	GetAll(ctx context.Context) ([]*model.Project, error)
	GetAllOverview(ctx context.Context) ([]*model.Project, error)
	UpdateName(ctx context.Context, uid, newName string) error
	Delete(ctx context.Context, uid string) error
	UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error
	//Relations
	AddDomain(ctx context.Context, projectUID, domainUID string) error
	AddTarget(ctx context.Context, projectUID, targetUID string) error
}

type DgraphProjectRepository struct {
	DB *dgo.Dgraph
}

func NewDgraphProjectRepository(db *dgo.Dgraph) *DgraphProjectRepository {
	return &DgraphProjectRepository{DB: db}
}

func (r *DgraphProjectRepository) Create(ctx context.Context, name string) (string, error) {
	txn := r.DB.NewTxn()
	defer txn.Discard(ctx)

	projectData := map[string]interface{}{
		"name":        name,
		"dgraph.type": "Project",
	}

	jsonData, err := json.Marshal(projectData)
	if err != nil {
		return "", fmt.Errorf("marshal error: %w", err)
	}

	mu := &api.Mutation{
		SetJson: jsonData,
	}

	assigned, err := txn.Mutate(ctx, mu)
	if err != nil {
		return "", fmt.Errorf("mutation error: %w", err)
	}

	if err := txn.Commit(ctx); err != nil {
		return "", fmt.Errorf("commit error: %w", err)
	}

	return assigned.Uids["blank-0"], nil
}

func (r *DgraphProjectRepository) Get(ctx context.Context, uid string) (*model.Project, error) {
	query := `
        query Project($uid: string) {
            project(func: uid($uid)) {
                uid
                name
                has_domain {
                    uid
                    name
                    has_host {
                        uid
                        ip
                        has_service {
                            uid
                            name
                            port
                        }
                    }
                    has_user {
                        uid
                        username
                        is_admin
                    }
                }
                has_target {
                    uid
                    ip_range
                }
            }
        }
    `

	vars := map[string]string{"$uid": uid}
	res, err := r.DB.NewTxn().QueryWithVars(ctx, query, vars)
	if err != nil {
		return nil, err
	}

	var result struct {
		Project []model.Project `json:"project"`
	}
	if err := json.Unmarshal(res.Json, &result); err != nil {
		return nil, err
	}

	if len(result.Project) == 0 {
		return nil, fmt.Errorf("project not found")
	}

	return &result.Project[0], nil
}

func (r *DgraphProjectRepository) GetAll(ctx context.Context) ([]*model.Project, error) {
	query := `
        {
            allProjects(func: type(Project)) {
                uid
                name
                has_domain {
                    uid
                    name
                    has_host {
                        uid
                        ip
                        has_service {
                            uid
                            name
                            port
                        }
                    }
                    has_user {
                        uid
                        username
                        is_admin
                    }
                }
                has_target {
                    uid
                    ip_range
                }
            }
        }
    `

	txn := r.DB.NewReadOnlyTxn()
	resp, err := txn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying Dgraph: %v", err)
	}

	var response struct {
		Projects []*model.Project `json:"allProjects"`
	}

	if err := json.Unmarshal(resp.Json, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return response.Projects, nil
}

func (r *DgraphProjectRepository) GetAllOverview(ctx context.Context) ([]*model.Project, error) {
	query := `
        {
            allProjects(func: type(Project)) {
                uid
                name
                type
            }
        }
    `

	txn := r.DB.NewReadOnlyTxn()
	resp, err := txn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error querying Dgraph: %v", err)
	}

	var response struct {
		Projects []*model.Project `json:"allProjects"`
	}

	if err := json.Unmarshal(resp.Json, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %v", err)
	}

	return response.Projects, nil
}

func (r *DgraphProjectRepository) UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error {
	fields["uid"] = uid

	updateJSON, err := json.Marshal(fields)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	mu := &api.Mutation{
		SetJson: updateJSON,
	}

	txn := r.DB.NewTxn()
	defer txn.Discard(ctx)

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("mutation error: %w", err)
	}

	return txn.Commit(ctx)
}

func (r *DgraphProjectRepository) Delete(ctx context.Context, uid string) error {
	txn := r.DB.NewTxn()
	defer txn.Discard(ctx)

	deleteProjectQuery := `{
		"uid": "` + uid + `",
		"name": null
	}`

	mutation := &api.Mutation{
		CommitNow:  true,
		DeleteJson: []byte(deleteProjectQuery),
	}

	_, err := txn.Mutate(context.Background(), mutation)

	return fmt.Errorf("error while deleting project in mutation: %w", err)
}

// TODO
func (r *DgraphProjectRepository) UpdateName(ctx context.Context, uid, newName string) error {
	panic("Not yet implemented")
}

func (r *DgraphProjectRepository) AddDomain(ctx context.Context, projectUID, domainUID string) error {
	txn := r.DB.NewTxn()
	defer txn.Discard(ctx)

	mu := &api.Mutation{
		SetNquads: []byte(fmt.Sprintf(
			`<%s> <has_domain> <%s> .`,
			projectUID, domainUID,
		)),
	}

	_, err := txn.Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("mutation error: %w", err)
	}

	return txn.Commit(ctx)
}

// TODO
func (r *DgraphProjectRepository) AddTarget(ctx context.Context, projectUID, targetUID string) error {
	panic("Not yet implemented")
}

/*// Service-Layer: Transaktion für atomare Erstellung
func (s *ProjectService) CreateFullProject(ctx context.Context, name string, domains []DomainConfig) (*model.Project, error) {
	txn := s.db.NewTxn()
	defer txn.Discard(ctx)

	// 1. Projekt erstellen
	projectUID, err := s.projectRepo.Create(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("project creation failed: %w", err)
	}

	// 2. Domains + Unterentitäten hinzufügen
	for _, domainCfg := range domains {
		domainUID, err := s.domainRepo.Create(ctx, domainCfg.Name)
		if err != nil {
			return nil, fmt.Errorf("domain creation failed: %w", err)
		}

		// Domain mit Projekt verknüpfen
		if err := s.projectRepo.AddDomain(ctx, projectUID, domainUID); err != nil {
			return nil, err
		}

		// Hosts zur Domain hinzufügen
		for _, hostIP := range domainCfg.Hosts {
			hostUID, err := s.hostRepo.Create(ctx, hostIP)
			if err != nil {
				return nil, err
			}
			if err := s.domainRepo.AddHost(ctx, domainUID, hostUID); err != nil {
				return nil, err
			}
		}
	}

	// 3. Transaktion commiten
	if err := txn.Commit(ctx); err != nil {
		return nil, fmt.Errorf("transaction failed: %w", err)
	}

	return s.projectRepo.GetFull(ctx, projectUID)
}


err := projectRepo.UpdateFields(ctx, "0x123", map[string]interface{}{
"description": "New project details",
"status":      "active",
})*/
