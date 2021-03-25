package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	router.Route("/items", func(router chi.Router) {
		router.Post("/", handleNewItems)
		router.Get("/", handleGetItems)
		router.Delete("/", handleDeleteItem)
	})

	http.ListenAndServe(":8080", router)
}

func handleNewItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "New items")
}

func handleGetItems(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get items")
}

func handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Delete item")
}
