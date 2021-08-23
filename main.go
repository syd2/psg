package psg

import (
	"context"
	"database/sql"
	"github.com/syd2/psg/db/db"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	conn, err := sql.Open("postgres", "username=postgres password=syd0101 dbname=psg sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	postgres := db.New(conn)
	log.Println("connected to the db")
	user := postgres.CreateUser(context.Background(), db.CreateUserParams{
		Username: "admin",
		Password: "admin0101",
	})
	log.Println(user)

}
