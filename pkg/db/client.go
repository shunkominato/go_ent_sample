package db

import (
	"context"
	"fmt"
	"go-gql-sample/app/ent"
	"go-gql-sample/app/ent/migrate"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Database struct {
	client *ent.Client
}

func GetDataSourceName() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?search_path=public&sslmode=disable",
	os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),  os.Getenv("DB_DB_NAME"))
}

func NewDatabase() (*Database, error) {
	client, err := ent.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DB_NAME")))
	if err != nil {
			return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Println("DB Connected")
	return &Database{client: client}, nil
}

func (d *Database) Close() {
  d.client.Close()
}
func (d *Database) EntClient() *ent.Client {
	return d.client
}

func (d *Database) Migrate() {
	err := d.client.Schema.Create(
			context.Background(),
			migrate.WithDropIndex(true),
			migrate.WithDropColumn(true),
	)
	if err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
	}
}