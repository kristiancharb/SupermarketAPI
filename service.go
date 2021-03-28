package main

import (
	"errors"
	"sync"
)

func getItems() *ItemList {
	ch := make(chan []Item)
	go fetch(ch)
	items := <-ch
	list := &ItemList{}
	list.Items = items
	return list
}

func getItem(code string) (*Item, error) {
	ch := make(chan *Item)
	go fetchItem(ch, code)
	item := <-ch
	if item == nil {
		return nil, &ItemNotFoundError{
			Code: code,
			Err:  errors.New("item not found"),
		}
	}
	return item, nil
}

func addItems(itemList *ItemList) []error {
	// Use WaitGroup to sync goroutines for each produce item
	wg := &sync.WaitGroup{}
	ch := make(chan error, len(itemList.Items))
	for _, item := range itemList.Items {
		wg.Add(1)
		go add(wg, ch, item)
	}
	wg.Wait()
	// Ensure we don't wait for more data when all routines finish
	close(ch)
	var errors []error
	for err := range ch {
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func deleteItem(code string) error {
	ch := make(chan error)
	go remove(ch, code)
	return <-ch
}
