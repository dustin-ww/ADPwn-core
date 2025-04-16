package repository

import (
	"context"
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

func (r *DraphTargetRepository) Create(ctx context.Context, tx *dgo.Txn, ip string, note string) (string, error) {
	// check if ip already exists with upsert
	query := fmt.Sprintf(`
        query checkIPRange($ip: string) {
            targets(func: eq(ip, "%s")) @filter(type(Target)) {
                v as uid
            }
        }
    `, ip)

	mu := &api.Mutation{
		SetNquads: []byte(fmt.Sprintf(`
            _:newTarget <ip> "%s" .
            _:newTarget <note> "%s" .
            _:newTarget <dgraph.type> "Target" .
        `, ip, note)),
		Cond: "@if(eq(len(v), 0))",
	}

	req := &api.Request{
		Query:     query,
		Mutations: []*api.Mutation{mu},
	}

	resp, err := tx.Do(ctx, req)
	if err != nil {
		return "", fmt.Errorf("upsert error: %w", err)
	}

	if len(resp.GetUids()) == 0 {
		return "", fmt.Errorf("ip_range '%s' already exists", ip)
	}

	return resp.GetUids()["newTarget"], nil
}

func (r *DraphTargetRepository) UpdateFields(ctx context.Context, uid string, fields map[string]interface{}) error {
	panic("implement me")
}
