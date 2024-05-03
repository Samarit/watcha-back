package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string) (pgx.Rows, error)
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

func (pg *postgres) Query(ctx context.Context, sql string) (pgx.Rows, error) {
	rows, err := pg.db.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return rows, nil
}
