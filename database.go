package main

import (
	"errors"
	"strings"
	"sync"
)

var itemsByCode map[string]*Item
var mutex *sync.RWMutex

func fetch() []Item {
	mutex.RLock()
	defer mutex.RUnlock()
	var items []Item
	for _, item := range itemsByCode {
		items = append(items, *item)
	}
	return items
}

func add(item Item) error {
	mutex.Lock()
	defer mutex.Unlock()
	item.Code = strings.ToUpper(item.Code)
	if _, exists := itemsByCode[item.Code]; exists {
		return &ItemConflictError{
			Code: item.Code,
			Err:  errors.New("item with code already exists"),
		}
	}
	itemsByCode[item.Code] = &item
	return nil
}

func remove(code string) error {
	mutex.Lock()
	defer mutex.Unlock()
	if !isValidCode(code) {
		return &BadCodeError{
			Code: code,
			Err:  errors.New("code is invalid"),
		}
	}
	if _, exists := itemsByCode[code]; !exists {
		return &ItemNotFoundError{
			Code: code,
			Err:  errors.New("item not found"),
		}
	}
	delete(itemsByCode, code)
	return nil
}

func initDb() {
	itemsByCode = map[string]*Item{
		"A12T-4GH7-QPL9-3N4M": {
			Code:  "A12T-4GH7-QPL9-3N4M",
			Name:  "Lettuce",
			Price: 3.46,
		},
		"E5T6-9UI3-TH15-QR88": {
			Code:  "E5T6-9UI3-TH15-QR88",
			Name:  "Peach",
			Price: 2.99,
		},
		"YRT6-72AS-K736-L4AR": {
			Code:  "YRT6-72AS-K736-L4AR",
			Name:  "Green Pepper",
			Price: 0.79,
		},
		"TQ4C-VV6T-75ZX-1RMR": {
			Code:  "TQ4C-VV6T-75ZX-1RMR",
			Name:  "Gala Apple",
			Price: 3.59,
		},
	}
	mutex = &sync.RWMutex{}
}

func init() {
	initDb()
}
