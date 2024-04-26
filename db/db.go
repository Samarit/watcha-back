package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	call (*sql.DB)
}

func DB() Database {
	db := Database{}

	return db
}
