package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/watcha-back/db"
	// "database/sql"
	// _ "github.com/lib/pq"
)

var host = "localhost:5555"

const (
	hostdb   = "localhost"
	portdb   = 5432
	user     = "postgres"
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

	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		hostdb, portdb, user, password, dbname)

	db, err := db.NewClient(context.TODO(), connString)
	if err != nil {
		log.Println("DB ERR")
	}

	db.GetUser(context.TODO())

	// rows, err := db.Exec("select * from users;")
	// if err != nil {
	// 	log.Println("SELECT ERR: ", err)
	// }
	// defer rows.Close()

	// user := User{}

	// for rows.Next() {
	// 	err := rows.Scan(&user.id, &user.login, &user.password)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(user)
	// }

	s.GET("/api", func(ctx *gin.Context) {
		fmt.Println("GET /")
		fmt.Println(user)
	})

	s.Run(host)

}
