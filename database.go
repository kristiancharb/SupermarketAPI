package main

var itemsByCode map[string]*Item

func fetch() []Item {
	var items []Item
	for _, item := range itemsByCode {
		items = append(items, *item)
	}
	return items
}

func add(item Item) {
	itemsByCode[item.Code] = &item
}

func remove(code string) {
	delete(itemsByCode, code)
}

func initDb() {
	itemsByCode = map[string]*Item{
		"A12T-4GH7-QPL9-3N4M": {
			Code: "A12T-4GH7-QPL9-3N4M",
			Name: "Lettuce",
			Price: 3.46,
		},
		"E5T6-9UI3-TH15-QR88": {
			Code: "E5T6-9UI3-TH15-QR88",
			Name: "Peach",
			Price: 2.99,
		},
		"YRT6-72AS-K736-L4AR": {
			Code: "YRT6-72AS-K736-L4AR",
			Name: "Green Pepper",
			Price: 0.79,
		},
		"TQ4C-VV6T-75ZX-1RMR": {
			Code: "TQ4C-VV6T-75ZX-1RMR",
			Name: "Gala Apple",
			Price: 3.59,
		},
	}
}

func init() {
	initDb()
}
