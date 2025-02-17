package main

import (
	"ADPwn/database/internal/utils"
	db_context "context"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"golang.org/x/net/context"
	"log"
	"os"
	"time"
)

func InitializeDB() {
	db, err := utils.GetDB()
	ctx, cancel := db_context.WithTimeout(db_context.Background(), 5*time.Second)

	defer cancel()

	schema, err := os.ReadFile("./adpwn.schema")
	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}
	log.Println("Schema file read successfully.")

	op := &api.Operation{
		Schema: string(schema),
	}

	err = db.Alter(context.Background(), op)

	if err != nil {
		log.Fatalf("Failed to read schema file: %v", err)
	}

	txn := db.NewTxn()
	defer txn.Discard(ctx)

}

func main() {
	InitializeDB()
}
