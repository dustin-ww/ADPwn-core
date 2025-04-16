package repository

import (
	"ADPwn-core/pkg/model"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

// ProjectRepository defines operations for project data access
type ProjectRepository interface {
	// CRUD
	Create(ctx context.Context, tx *dgo.Txn, name string) (string, error)
	Get(ctx context.Context, tx *dgo.Txn, uid string) (*model.Project, error)
	GetAll(ctx context.Context, tx *dgo.Txn) ([]*model.Project, error)
	// TODO: Move into target repo
	GetTargets(ctx context.Context, tx *dgo.Txn, uid string) ([]*model.Target, error)
	Delete(ctx context.Context, tx *dgo.Txn, uid string) error
	UpdateFields(ctx context.Context, tx *dgo.Txn, uid string, fields map[string]interface{}) error

	// Relations
	AddDomain(ctx context.Context, tx *dgo.Txn, projectUID, domainUID string) error
	AddTarget(ctx context.Context, tx *dgo.Txn, projectUID, targetUID string) error
}

// DgraphProjectRepository implements ProjectRepository using Dgraph
type DgraphProjectRepository struct{}

// NewDgraphProjectRepository creates a new Dgraph project repository
func NewDgraphProjectRepository() *DgraphProjectRepository {
	return &DgraphProjectRepository{}
}

// Create adds a new project to the database
func (r *DgraphProjectRepository) Create(ctx context.Context, tx *dgo.Txn, name string) (string, error) {
	projectData := map[string]interface{}{
		"name":        name,
		"dgraph.type": "Project",
		"created_at":  time.Now(),
	}

	jsonData, err := json.Marshal(projectData)
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

// Get retrieves a project by UID
func (r *DgraphProjectRepository) Get(ctx context.Context, tx *dgo.Txn, uid string) (*model.Project, error) {
	query := `
        query Project($uid: string) {
            project(func: uid($uid)) {
                uid
                name
                tags
                created_at
                modified_at
                description
                has_domain {
                    uid
                }
                has_target {
                    uid
                    ip_range
                    name
                }
            }
        }
    `

	vars := map[string]string{"$uid": uid}
	res, err := tx.QueryWithVars(ctx, query, vars)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	var result struct {
		Project []model.Project `json:"project"`
	}
	if err := json.Unmarshal(res.Json, &result); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	if len(result.Project) == 0 {
		return nil, fmt.Errorf("project not found: %s", uid)
	}

	return &result.Project[0], nil
}

// GetTargets retrieves all targets for a project
func (r *DgraphProjectRepository) GetTargets(ctx context.Context, tx *dgo.Txn, uid string) ([]*model.Target, error) {
	query := `
        query Project($uid: string) {
            project(func: uid($uid)) @filter(eq(dgraph.type, "Project")) {
                uid
                has_target {
                    uid
                    ip_range
                    name
                    dgraph.type
                }
            }
        }
    `

	vars := map[string]string{"$uid": uid}
	res, err := tx.QueryWithVars(ctx, query, vars)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	var result struct {
		Project []struct {
			UID       string          `json:"uid"`
			HasTarget []*model.Target `json:"has_target"`
		} `json:"project"`
	}

	if err := json.Unmarshal(res.Json, &result); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	if len(result.Project) == 0 {
		return nil, fmt.Errorf("project not found: %s", uid)
	}

	targets := result.Project[0].HasTarget

	if targets == nil {
		return []*model.Target{}, nil
	}

	return targets, nil
}

// GetAll retrieves all projects with full details
func (r *DgraphProjectRepository) GetAll(ctx context.Context, tx *dgo.Txn) ([]*model.Project, error) {
	query := `
        {
            allProjects(func: type(Project)) {
                uid
                name
				type
                description
                modified_at
                created_at
                tags
                has_domain {
                    uid
                }
                has_target {
                    uid
                    ip_range
                }
            }
        }
    `

	resp, err := tx.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}

	var response struct {
		Projects []*model.Project `json:"allProjects"`
	}

	if err := json.Unmarshal(resp.Json, &response); err != nil {
		return nil, fmt.Errorf("unmarshal error: %w", err)
	}

	return response.Projects, nil
}

// UpdateFields updates specified fields on a project
func (r *DgraphProjectRepository) UpdateFields(ctx context.Context, tx *dgo.Txn, uid string, fields map[string]interface{}) error {
	fields["uid"] = uid
	fields["updated_at"] = time.Now().Format(time.RFC3339)

	updateJSON, err := json.Marshal(fields)
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	mu := &api.Mutation{
		SetJson: updateJSON,
	}

	_, err = tx.Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("mutation error: %w", err)
	}
	return nil
}

// Delete removes a project by UID
func (r *DgraphProjectRepository) Delete(ctx context.Context, tx *dgo.Txn, uid string) error {
	deleteQuery := fmt.Sprintf(`{"uid": "%s", "name": null}`, uid)

	mutation := &api.Mutation{
		DeleteJson: []byte(deleteQuery),
	}

	_, err := tx.Mutate(ctx, mutation)
	if err != nil {
		return fmt.Errorf("mutation error: %w", err)
	}

	return nil
}

// AddDomain connects a domain to a project
func (r *DgraphProjectRepository) AddDomain(ctx context.Context, tx *dgo.Txn, projectUID, domainUID string) error {
	nquad := fmt.Sprintf(`<%s> <has_domain> <%s> .`, projectUID, domainUID)

	mu := &api.Mutation{
		SetNquads: []byte(nquad),
	}

	_, err := tx.Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("mutation error: %w", err)
	}

	return nil
}

// AddTarget connects a target to a project
func (r *DgraphProjectRepository) AddTarget(ctx context.Context, tx *dgo.Txn, projectUID, targetUID string) error {
	nquad := fmt.Sprintf(`<%s> <has_target> <%s> .`, projectUID, targetUID)

	mu := &api.Mutation{
		SetNquads: []byte(nquad),
	}

	_, err := tx.Mutate(ctx, mu)
	if err != nil {
		return fmt.Errorf("mutation error: %w", err)
	}

	return nil
}
