package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	router.Route("/items", func(router chi.Router) {
		router.Post("/", handleNewItems)
		router.Get("/", handleGetItems)
		router.Delete("/{code}", handleDeleteItem)
	})

	http.ListenAndServe(":8080", router)
}

func handleNewItems(w http.ResponseWriter, r *http.Request) {
	items := &ItemList{}
	render.Bind(r, items)
	addItems(items)
}

func handleGetItems(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, getItems())
}

func handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	deleteItem(code)
}
