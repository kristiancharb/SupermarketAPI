package main

import (
	"errors"
	"strings"
	"sync"
)

// Global item store
var itemsByCode map[string]*Item

// Sync access to global item store
var mutex *sync.RWMutex

func fetch(ch chan []Item) {
	mutex.RLock()
	defer mutex.RUnlock()
	var items []Item
	for _, item := range itemsByCode {
		items = append(items, *item)
	}
	ch <- items
}

func fetchItem(ch chan *Item, code string) {
	mutex.RLock()
	defer mutex.RUnlock()
	code = strings.ToUpper(code)
	item, exists := itemsByCode[code]
	if !exists {
		ch <- nil
		return
	}
	ch <- item
}

func add(wg *sync.WaitGroup, ch chan error, item Item) {
	mutex.Lock()
	defer mutex.Unlock()
	defer wg.Done()
	item.Code = strings.ToUpper(item.Code)
	if _, exists := itemsByCode[item.Code]; exists {
		ch <- &ItemConflictError{
			Code: item.Code,
			Err:  errors.New("item with code already exists"),
		}
		return
	}
	itemsByCode[item.Code] = &item
}

func remove(ch chan error, code string) {
	mutex.Lock()
	defer mutex.Unlock()
	code = strings.ToUpper(code)
	if !isValidCode(code) {
		ch <- &BadCodeError{
			Code: code,
			Err:  errors.New("code is invalid"),
		}
		return
	}
	if _, exists := itemsByCode[code]; !exists {
		ch <- &ItemNotFoundError{
			Code: code,
			Err:  errors.New("item not found"),
		}
		return
	}
	delete(itemsByCode, code)
	ch <- nil
}

func initDb() {
	// Add initial produce items to item store
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
