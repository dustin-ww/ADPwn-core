package utils

import (
	"log"
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

		dialOpts := append([]grpc.DialOption{},
			grpc.WithInsecure(),
			grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)),
		)
		var err error

		conn, err := grpc.Dial("localhost:9080", dialOpts...)
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
