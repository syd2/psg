package main

import (
	"context"
	"database/sql"
	"log"

	"github.com/syd2/psg/db/db"

	_ "github.com/lib/pq"
)

func main() {
	conn, err := sql.Open("postgres", "user=postgres password=syd0101 dbname=psg sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	postgres := db.New(conn)
	log.Println("connected to the db")

	//creating new user
	newUser(postgres, "admin", "admin0101")

}

func newUser(postgres *db.Queries, username string, password string) {
	error := postgres.CreateUser(context.Background(), db.CreateUserParams{
		Username: username,
		Password: password,
	})
	if error != nil {
		log.Fatal(error)
	}
	log.Println("Created")
}
