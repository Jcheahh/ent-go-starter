package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"entdemo/ent"
	"entdemo/ent/user"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
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

func CreateCarOwner(ctx context.Context, u *ent.User, c *ent.Car) error {
	car, err := c.Update().SetOwner(u).Save(ctx)
	if err != nil {
		return fmt.Errorf("failed updating car owner: %w", err)
	}
	log.Printf("car %q owner was updated to %q, ", car.Model, u.Name)
	return nil
}

func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
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
	client := Open("postgresql://postgres:mysecretpassword@127.0.0.1/newdb")

	// Your code. For example:
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatal(err)
	}

	// Create a new user.
	// createdUser, err := CreateUser(ctx, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("CreatedUser: ", createdUser.ID)
	// users, err := client.User.Query().All(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(users)
	// newUserWithCar, err := CreateCars(ctx, client)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("newUserWithCar: ", newUserWithCar.ID)
	// Get a user by ID
	u, err := client.User.Get(ctx, 2)
	if err != nil {
		log.Fatal(err)
	}
	// Get a car by ID
	// c, err := client.Car.Get(ctx, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	cars, _ := client.User.QueryCars(u).All(ctx)
	fmt.Println("cars: ", cars)

	toyota, err := client.Car.
		Create().
		SetModel("toyota").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		fmt.Println("failed creating car: %w", err)
	}
	log.Println("car was created: ", toyota)

	setCarOwnerErr := CreateCarOwner(ctx, u, toyota)
	if setCarOwnerErr != nil {
		log.Fatal(setCarOwnerErr)
	}
}
