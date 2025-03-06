package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

type TargetRepository interface {
	//CRUD
	Create(ctx context.Context, ipRange string) (string, error) // Returns UID
	UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error
}

type DraphTargetRepository struct {
	DB *dgo.Dgraph
}

func NewDgraphTargetRepository(db *dgo.Dgraph) *DraphTargetRepository {
	return &DraphTargetRepository{DB: db}
}

func (r *DraphTargetRepository) Create(ctx context.Context, ipRange string) (string, error) {
	txn := r.DB.NewTxn()
	defer txn.Discard(ctx)

	projectData := map[string]interface{}{
		"ip_range":    ipRange,
		"dgraph.type": "Target",
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

func (r *DraphTargetRepository) UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error {
	panic("implement me")
}
