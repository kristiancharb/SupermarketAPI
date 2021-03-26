package main

import (
	"fmt"
	"net/http"
	"regexp"
)

type Item struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price,omitempty"`
}

type ItemList struct {
	Items []Item `json:"items"`
}

type BadCodeError struct {
	Code string
	Err  error
}

type ItemNotFoundError struct {
	Code string
	Err  error
}

type ItemConflictError struct {
	Code string
	Err  error
}

func (item *Item) Bind(r *http.Request) error {
	if item.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	if item.Code == "" {
		return fmt.Errorf("code is a required field")
	}
	if item.Price == 0 {
		return fmt.Errorf("price is a required field")
	}
	if !isValidCode(item.Code) {
		return fmt.Errorf("code is invalid")
	}
	return nil
}

func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (item *Item) equals(other *Item) bool {
	return (item.Name == other.Name &&
		item.Code == other.Code &&
		item.Price == other.Price)
}

func (*ItemList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (list *ItemList) Bind(r *http.Request) error {
	for _, item := range list.Items {
		err := item.Bind(r)
		if err != nil {
			return err
		}
	}
	return nil
}

func isValidCode(code string) bool {
	codePattern := "^(?:[A-Za-z0-9]{4}-){3}[A-Za-z0-9]{4}$"
	matched, _ := regexp.MatchString(codePattern, code)
	return matched
}

func (e *BadCodeError) Error() string {
	return e.Code + ": " + e.Err.Error()
}

func (e *ItemNotFoundError) Error() string {
	return e.Code + ": " + e.Err.Error()
}

func (e *ItemConflictError) Error() string {
	return e.Code + ": " + e.Err.Error()
}

