package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

type ErrorResponse struct {
	Message    string `json:"message"`
	StatusCode int    `json:"-"`
}

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
	err := render.Bind(r, items)
	if err != nil {
		render.Render(w, r, &ErrorResponse{
			Message:    err.Error(),
			StatusCode: 400,
		})
		return
	}
	err = addItems(items)
	if err != nil {
		render.Render(w, r, &ErrorResponse{
			Message:    err.Error(),
			StatusCode: 409,
		})
	}
}

func handleGetItems(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, getItems())
}

func handleDeleteItem(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	err := deleteItem(code)
	if err != nil {
		switch err.(type) {
		case *BadCodeError:
			render.Render(w, r, &ErrorResponse{
				Message:    err.Error(),
				StatusCode: 400,
			})
		default:
			render.Render(w, r, &ErrorResponse{
				Message:    err.Error(),
				StatusCode: 404,
			})
		}
	}
}

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}
