package main

import (
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"github.com/syd2/psg/api"
	"github.com/syd2/psg/api/handlers"
	"github.com/syd2/psg/server"
)

func main() {

	router := chi.NewRouter()
	s := server.NewServer()

	router.Group(func(r chi.Router) {
		r.Get("/", handlers.HelloWorldHandler(s.Queries))
		r.Post("/users/create", handlers.CreateUserHandler(s.Queries, "syd0101"))
		r.Post("/users/login", handlers.LoginUserHandler(s.Queries, "syd0101"))
	})

	//private routes
	router.Group(func(r chi.Router) {
		r.Use(api.AuthWrapper("syd0101"))
		r.Get("/test", handlers.HelloWorldHandler(s.Queries))
		r.Post("create/password", handlers.CreatePasswordHandler(s.Queries, "syd0101"))

	})
	s.Router.Mount("/", router)
	s.RunServer()

}
