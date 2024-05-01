package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

type Client interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

type postgres struct {
	db *pgxpool.Pool
}

var instance *postgres
var once sync.Once

func NewClient(ctx context.Context, connString string) (*postgres, error) {
	once.Do(func() {
		db, err := pgxpool.New(ctx, connString)
		if err != nil {
			log.Fatal("unable to create connection pool: %w", err)
		}

		instance = &postgres{db}
	})

	return instance, nil
}

func (pg *postgres) GetUser(ctx context.Context) error {
	query := `SELECT * FROM USERS;`

	result, err := pg.db.Exec(ctx, query)
	if err != nil {
		fmt.Println("QUERY ERROR")
	}

	fmt.Println(result.String())

	return nil
}
