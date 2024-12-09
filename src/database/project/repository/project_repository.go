package repository

import (
	"ADPwn/database/project/model"
	"context"
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

type ProjectRepository interface {
	AllProjects(ctx context.Context) ([]model.Project, error)
	SaveProject(ctx context.Context, project model.Project) error
	DeleteProject(ctx context.Context, project model.Project) error
	AllConnectedUIDs(ctx context.Context, project model.Project) ([]string, error)
}

type DgraphIOProjectRepository struct {
	DB *dgo.Dgraph
}

func NewDgraphIOProjectRepository(db *dgo.Dgraph) *DgraphIOProjectRepository {
	return &DgraphIOProjectRepository{DB: db}
}

func (r *DgraphIOProjectRepository) AllProjects(ctx context.Context) ([]model.Project, error) {

	txn := r.DB.NewTxn()
	defer txn.Discard(ctx)

	query := `{
		allProjects(func: has(name)) {
    		uid
			name
    		hosts {
				uid
        		ip
        		host_project_id
        		is_domaincontroller
      	}
  	}
}`

	res, err := txn.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("error while querying dgraph: %v", err)
	}

	var response struct {
		AllProjects []model.Project `json:"allProjects"`
	}

	if err := json.Unmarshal(res.Json, &response); err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %v", err)
	}

	return response.AllProjects, nil
}

func (r *DgraphIOProjectRepository) SaveProject(ctx context.Context, project model.Project) error {
	txn := r.DB.NewTxn()
	defer txn.Discard(ctx)

	pj, err := json.Marshal(project)
	if err != nil {
		return err
	}

	mu := &api.Mutation{SetJson: pj}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	txn.Commit(ctx)
	fmt.Println("Saved Project")

	return err
}

func (r *DgraphIOProjectRepository) DeleteProject(ctx context.Context, project model.Project) error {

	txn := r.DB.NewTxn()
	defer txn.Discard(ctx)

	query := `{
		allUIDs(func: eq(uid, "` + project.UID + `")) {
		  uid
		  hosts {
			uid
		  }
		}
	  }`

	res, err := txn.Query(ctx, query)
	if err != nil {
		return fmt.Errorf("error while querying dgraph: %v", err)
	}

	var response struct {
		AllUids []string `json:"allUIDs"`
	}

	if err := json.Unmarshal(res.Json, &response); err != nil {
		return fmt.Errorf("error unmarshaling json: %v", err)
	}

	return nil
}

func (r *DgraphIOProjectRepository) AllConnectedUIDs(ctx context.Context, project model.Project) ([]string, error) {
	return nil, nil
}

/* func (r *DgraphIOProjectRepository) SaveHosts(ctx context.Context, project model.Project, hosts []model.Host) error {
	txn := r.DB.NewTxn()
	defer txn.Discard(ctx)

	projectUpdate := Project{
		UID:   project.UID, // UID des bestehenden Projekts
		Hosts: hosts, // Die neuen Clients
	}


	pj, err := json.Marshal(hosts)
	if err != nil {
		return err
	}

	mu := &api.Mutation{SetJson: pj}

	_, err = txn.Mutate(ctx, mu)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	txn.Commit(ctx)

	return nil
} */
