package main

import (
	"context"
	"fmt"
	"go-gql-sample/app/ent"
	"go-gql-sample/app/ent/migrate"
	"go-gql-sample/app/ent/user"
	"go-gql-sample/app/pkg/config"
	"go-gql-sample/app/pkg/db"
	"log"
	"os"
	"time"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
// 	u, err := client.User.
// 			Create().
// 			SetAge(30).
// 			SetName("a8m").
// 			Save(ctx)
// 	if err != nil {
// 			return nil, fmt.Errorf("failed creating user: %w", err)
// 	}
// 	log.Println("user was created: ", u)
// 	return u, nil
// }

// func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
// 	u, err := client.User.
// 			Query().
// 			Where(user.Name("a8m")).
// 			// `Only` fails if no user found,
// 			// or more than 1 user returned.
// 			Only(ctx)
// 	if err != nil {
// 			return nil, fmt.Errorf("failed querying user: %w", err)
// 	}
// 	log.Println("user returned: ", u)
// 	return u, nil
// }


func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
			Create().
			SetAge(30).
			SetName("a8m").
			Save(ctx)
	if err != nil {
			return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
			Query().
			Where(user.Name("a8m")).
			// `Only` fails if no user found,
			// or more than 1 user returned.
			Only(ctx)
	if err != nil {
			return nil, fmt.Errorf("failed querying user: %w", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// Create a new car with model "Tesla".
	tesla, err := client.Car.
			Create().
			SetModel("Tesla").
			SetRegisteredAt(time.Now()).
			Save(ctx)
	if err != nil {
			return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", tesla)

	// Create a new car with model "Ford".
	ford, err := client.Car.
			Create().
			SetModel("Ford").
			SetRegisteredAt(time.Now()).
			Save(ctx)
	if err != nil {
			return nil, fmt.Errorf("failed creating car: %w", err)
	}
	log.Println("car was created: ", ford)

	// Create a new user, and add it the 2 cars.
	a8m, err := client.User.
			Create().
			SetAge(30).
			SetName("a8m").
			AddCars(tesla, ford).
			Save(ctx)
	if err != nil {
			return nil, fmt.Errorf("failed creating user: %w", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}

func QueryCarUsers(ctx context.Context, client *ent.Client) error {
	u, err := client.User.
	Query().
	Where(user.Name("a8m")).
	// `Only` fails if no user found,
	// or more than 1 user returned.
	Only(ctx)

	cars, err := u.QueryCars().All(ctx)
	if err != nil {
			return fmt.Errorf("failed querying user cars: %w", err)
	}
	// Query the inverse edge.
	for _, c := range cars {
			owner, err := c.QueryOwner().Only(ctx)
			if err != nil {
					return fmt.Errorf("failed querying car %q owner: %w", c.Model, err)
			}
			log.Printf("car %q owner: %q\n", c.Model, owner.Name)
	}
	return nil
}

func main() {
	godotenv.Load(".env")

	config.SetConfig()
	// db, _ := db.NewDatabase()	
	// client := db.EntClient()
	// defer client.Close()

	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	log.Fatalf("failed creating schema resources: %v", err)
	// }
	// QueryCarUsers(context.Background(), client)

	ctx := context.Background()

	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}

    // Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeInspect), // provide migration mode
		schema.WithDialect(dialect.Postgres),           // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}

	dataSourceName := db.GetDataSourceName()
	// dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DB_NAME"))
	// dataSourceName := fmt.Sprintf("postgres://postgres:postgres@gql-db-1:5432/engin?search_path=public&sslmode=disable")

	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, dataSourceName, os.Args[1], opts...)
	if err != nil {
			log.Fatalf("failed generating migration file: %v", err)
	}
}