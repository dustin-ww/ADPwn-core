package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

var (
	db    *dgo.Dgraph
	once  sync.Once
	dbErr error
)

func GetDB() (*dgo.Dgraph, error) {
	once.Do(func() {

		// Connection String
		host := os.Getenv("DGRAPH_HOST")
		port := os.Getenv("DGRAPH_PORT")

		dialOpts := append([]grpc.DialOption{},
			grpc.WithInsecure(),
			grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)),
		)
		var err error

		if host == "" {
			host = "localhost"
			log.Println("WARNING: Dgraph host not set, using localhost")
		}

		conn, err := grpc.NewClient(host+port, dialOpts...)
		if err != nil {
			dbErr = err
			log.Fatal(err)
		}

		db = dgo.NewDgraphClient(api.NewDgraphClient(conn))
	})

	if dbErr != nil {
		return nil, dbErr
	}

	return db, nil
}

func ExecuteInTransaction(ctx context.Context, db *dgo.Dgraph, op func(tx *dgo.Txn) error) error {
	tx := db.NewTxn()
	defer tx.Discard(ctx)

	if err := op(tx); err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit failed: %w", err)
	}
	return nil
}

func ExecuteRead[T any](ctx context.Context, db *dgo.Dgraph, op func(tx *dgo.Txn) (T, error)) (T, error) {
	tx := db.NewReadOnlyTxn()
	defer tx.Discard(ctx)
	return op(tx)
}
