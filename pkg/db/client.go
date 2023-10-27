package db

import (
	"fmt"
	"go-gql-sample/app/ent"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	client *ent.Client
}

func NewDatabase() (*Database, error) {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DB_NAME")))
	if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &Database{client: client}, nil
}

func (d *Database) Close() {
  d.client.Close()
}
func (d *Database) EntClient() *ent.Client {
	return d.client
}