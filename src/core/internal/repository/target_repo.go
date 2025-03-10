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
	Create(ctx context.Context, tx *dgo.Txn, ipRange string, name string) (string, error) // Returns UID
	UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error
}

type DraphTargetRepository struct {
	DB *dgo.Dgraph
}

func NewDgraphTargetRepository(db *dgo.Dgraph) *DraphTargetRepository {
	return &DraphTargetRepository{DB: db}
}

func (r *DraphTargetRepository) Create(ctx context.Context, tx *dgo.Txn, ipRange string, name string) (string, error) {
	// Füge eine explizite Blank Node UID hinzu
	target := map[string]interface{}{
		"uid":         "_:newTarget", // Wichtig: Blank Node Identifier
		"name":        name,
		"ip_range":    ipRange,
		"dgraph.type": "Target",
	}

	targetJson, err := json.Marshal(target)
	if err != nil {
		return "", fmt.Errorf("marshal error: %w", err)
	}

	mu := &api.Mutation{
		SetJson: targetJson,
	}
	assigned, err := tx.Mutate(ctx, mu)
	if err != nil {
		return "", fmt.Errorf("mutation error: %w", err)
	}

	return assigned.Uids["newTarget"], nil
}

func (r *DraphTargetRepository) UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error {
	panic("implement me")
}
