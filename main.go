package main

import (
	"log"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/syd2/psg/api"
	"github.com/syd2/psg/api/handlers"
	"github.com/syd2/psg/server"
	"github.com/syd2/psg/utils"
)

func main() {

	router := chi.NewRouter()
	s := server.NewServer()
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("Cant log config")
	}
	secret := config.Secret
	router.Group(func(r chi.Router) {
		r.Get("/", handlers.HelloWorldHandler(s.Queries))
		r.Post("/users/create", handlers.CreateUserHandler(s.Queries, secret))
		r.Post("/users/login", handlers.LoginUserHandler(s.Queries, secret))
	})

	//private routes
	router.Group(func(r chi.Router) {
		r.Use(api.AuthWrapper(secret))
		r.Get("/test", handlers.HelloWorldHandler(s.Queries))
		r.Post("/create/password", handlers.CreatePasswordHandler(s.Queries, secret))

	})
	s.Router.Mount("/", router)
	s.RunServer()

}
