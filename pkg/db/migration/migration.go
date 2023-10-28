package migration

import (
	"context"
	"go-gql-sample/app/pkg/db"
	"log"

	_ "github.com/lib/pq"
)
func main() {
	db, _ := db.NewDatabase()	
	client := db.EntClient()

	defer db.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}