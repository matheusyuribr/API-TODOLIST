package main

import (
	"api-postgresql/configs"
	"api-postgresql/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", handlers.Create)

	r.Get("/{id}", handlers.Get)

	r.Delete("/{id}", handlers.Delete)

	r.Get("/", handlers.List)

	r.Put("/{id}", handlers.Update)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)

}
