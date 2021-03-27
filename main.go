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

type MultipleErrorResponse struct {
	Errors     []string `json:"errors"`
	StatusCode int      `json:"-"`
}

func main() {
	router := chi.NewRouter()

	router.Use(middleware.Recoverer)

	router.Route("/items", func(router chi.Router) {
		router.Post("/", handleNewItems)
		router.Get("/", handleGetItems)
		router.Get("/{code}", handleGetItem)
		router.Delete("/{code}", handleDeleteItem)
	})

	http.ListenAndServe(":8080", router)
}

func handleNewItems(w http.ResponseWriter, r *http.Request) {
	itemList := &ItemList{}
	err := render.Bind(r, itemList)
	if err != nil {
		render.Render(w, r, &ErrorResponse{
			Message:    err.Error(),
			StatusCode: 400,
		})
		return
	}
	errors := addItems(itemList)
	statusCode := 200
	if len(errors) == len(itemList.Items) {
		statusCode = 409
	}
	if len(errors) > 0 {
		render.Render(w, r, &MultipleErrorResponse{
			Errors:     getErrorMessages(errors),
			StatusCode: statusCode,
		})
	}
}

func handleGetItems(w http.ResponseWriter, r *http.Request) {
	render.Render(w, r, getItems())
}

func handleGetItem(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	item, err := getItem(code)
	if err != nil {
		render.Render(w, r, &ErrorResponse{
			Message:    err.Error(),
			StatusCode: 404,
		})
		return
	}
	render.Render(w, r, item)
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

func (e *MultipleErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

func getErrorMessages(errors []error) []string {
	var messages []string
	for _, err := range errors {
		messages = append(messages, err.Error())
	}
	return messages
}
