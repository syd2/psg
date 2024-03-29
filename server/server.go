package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/syd2/psg/db/db"
	"github.com/syd2/psg/utils"
)

type Server struct {
	Router  *chi.Mux
	Queries *db.Queries
}

func NewServer() *Server {
	s := &Server{}

	//loading configs
	config, err := utils.LoadConfig(".")

	//preparing the db
	dbSource := config.DBSource
	dbDriver := config.DBDriver
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the db : ", err.Error())

	}
	s.Queries = db.New(conn)
	log.Println("Connected to the database")

	//preparing the router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.Recoverer)

	s.Router = r
	return s

}

func (s *Server) RunServer() {

	fmt.Println("Serving on port 3000....")
	http.ListenAndServe(":3000", s.Router)
}
