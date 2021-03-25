package main

import (
	"fmt"
	"net/http"
)

type Item struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price,omitempty"`
}

type ItemList struct {
	Items []Item `json:"items"`
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



