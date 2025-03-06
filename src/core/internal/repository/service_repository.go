package repository

import (
	"ADPwn/core/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

type ServiceRepository interface {
	//CRUD
	Create(ctx context.Context, name string) (string, error) // Returns UID
	CreateWithObject(ctx context.Context, model model.Service) (string, error)
	Get(ctx context.Context, uid string) (*model.Service, error)
	UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error
	//Relations
}

type DgraphServiceRepository struct {
	DB *dgo.Dgraph
}

func (r *DgraphServiceRepository) UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error {
	//TODO implement me
	panic("implement me")
}

func NewDgraphServiceRepository(db *dgo.Dgraph) *DgraphServiceRepository {
	return &DgraphServiceRepository{DB: db}
}

func (r *DgraphServiceRepository) Create(ctx context.Context, name string) (string, error) {
	service := &model.Service{}
	service.Name = name
	return r.CreateWithObject(ctx, *service)
}

func (r *DgraphServiceRepository) CreateWithObject(ctx context.Context, service model.Service) (string, error) {
	txn := r.DB.NewTxn()
	defer txn.Discard(ctx)

	jsonData, err := json.Marshal(service)
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

func (r *DgraphServiceRepository) Get(ctx context.Context, uid string) (*model.Service, error) {
	panic("implement me")
}
