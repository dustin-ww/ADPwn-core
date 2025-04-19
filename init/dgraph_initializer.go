package init

import (
	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"golang.org/x/net/context"
	"log"
	"os"
	"time"
)

func InitializeDgraphSchema(db *dgo.Dgraph) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	schema, err := os.ReadFile("./adpwn.schema")
	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}
	log.Println("Dgraph: Schema file read successfully.")

	op := &api.Operation{
		Schema: string(schema),
	}

	err = db.Alter(context.Background(), op)
	if err != nil {
		log.Fatalf("Dgraph: Failed to alter schema: %v", err)
	}

	txn := db.NewTxn()
	defer txn.Discard(ctx)
}

/*
func main() {
	InitializeDB()
	DropAllPostgresObjects()
	InitializePostgresDB()
}
*/
