package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var host = "localhost:5555"

const (
	hostdb   = "localhost"
	portdb   = 5432
	user     = "pguser"
	password = "pgpass"
	dbname   = "postgres"
)

type User struct {
	id       int
	login    string
	password string
}

func main() {
	s := gin.Default()

	db, err := initDB()
	if err != nil {
		log.Println("DB ERR")
	}

	rows, err := db.Query("select * from users;")
	if err != nil {
		log.Println("SELECT ERR")
	}
	defer rows.Close()

	user := User{}

	for rows.Next() {
		err := rows.Scan(&user.id, &user.login, &user.password)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(user)
	}

	s.GET("/api", func(ctx *gin.Context) {
		fmt.Println("GET /")
		fmt.Println(user)
	})

	s.Run(host)

}

func initDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostdb, portdb, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil
}

func getUser() User {
	user := User{}

	return user
}
