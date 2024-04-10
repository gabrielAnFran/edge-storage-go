package main

import (
	"fmt"
	"net/http"

	"github.com/gabrielAnFran/edge-storage-go/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {

	fmt.Println("Server running...")

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/buckets", handlers.ListBuckets)
	r.Delete("/buckets/{name}", handlers.DeleteBucket)
	r.Post("/buckets", handlers.CreateBucket)

	http.ListenAndServe(":8080", r)

}
