package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
)

type HostRepository interface {
	Create(ctx context.Context, ip string) (string, error)
	SetDomainController(ctx context.Context, hostUID string, isDC bool) error
	AddService(ctx context.Context, hostUID, serviceUID string) error
}

type DraphHostRepository struct {
	DB *dgo.Dgraph
}

func NewDgraphHostRepository(db *dgo.Dgraph) *DraphHostRepository {
	return &DraphHostRepository{DB: db}
}

func (r *DraphHostRepository) Create(ctx context.Context, ip string) (string, error) {
	txn := r.DB.NewTxn()
	defer txn.Discard(ctx)

	hostData := map[string]interface{}{
		"ip":          ip,
		"dgraph.type": "Host",
	}

	jsonData, err := json.Marshal(hostData)
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

// TODO
func (r *DraphHostRepository) SetDomainController(ctx context.Context, hostUID string, isDC bool) error {
	return nil
}

// TODO
func (r *DraphHostRepository) AddService(ctx context.Context, hostUID, serviceUID string) error {
	return nil
}
