package main

import (
	"sort"
	"testing"
)

func TestGetItems(t *testing.T) {
	initDb()
	expected := getInitialTestItems()
	actual := getItems().Items
	compareItemsLists(t, expected, actual)
}

func TestAddItems(t *testing.T) {
	initDb()
	list := &ItemList{
		Items: []Item{
			{
				Code:  "HJ23-823K-K82K-203L",
				Name:  "Ice Cream",
				Price: 5.99,
			},
			{
				Code:  "8JKL-82JH-23LK-93LC",
				Name:  "Bread",
				Price: 1.29,
			},
		},
	}
	addItems(list)
	expected := getInitialTestItems()
	expected = append(expected, list.Items...)
	compareItemsLists(t, expected, fetch())
}

func TestDeleteItem(t *testing.T) {
	initDb()
	expected := getInitialTestItems()
	deleteItem("A12T-4GH7-QPL9-3N4M")
	compareItemsLists(t, expected[1:], fetch())
}

func TestIsValid(t *testing.T) {
	validCodes := []string{
		"A12T-4GH7-QPL9-3N4M",
		"A12t-4Gh7-QPl9-3N4m",
		"1111-1111-1111-1111",
		"HHHH-HHHH-HHHH-HHHH",
	}
	invalidCodes := []string{
		"A12-4GH7-QPL9-3N4M",
		"A12T4GH7-QPL9-3N4M",
		"A12T-4GH7-QPL9-3N4MS",
		"A12?-4GH7-QPL9-3N4MS",
	}
	checkCodes(t, validCodes, true)
	checkCodes(t, invalidCodes, false)
}

func compareItemsLists(t *testing.T, expected, actual []Item) {
	sort.SliceStable(expected, func(i, j int) bool {
		return expected[i].Code < expected[j].Code
	})
	sort.SliceStable(actual, func(i, j int) bool {
		return actual[i].Code < actual[j].Code
	})
	for i := range expected {
		if !expected[i].equals(&actual[i]) {
			t.Fatalf(`Expected: %+v Actual: %+v`, expected[i], actual[i])
		}
	}
}

func checkCodes(t *testing.T, codes []string, expected bool) {
	for _, code := range codes {
		if isValid := isValidCode(code); isValid != expected {
			t.Fatalf(`%s -> Expected: %t Actual: %t`, code, expected, isValid)
		}
	}
}

func getInitialTestItems() []Item {
	items := []Item{
		{
			Code:  "A12T-4GH7-QPL9-3N4M",
			Name:  "Lettuce",
			Price: 3.46,
		},
		{
			Code:  "E5T6-9UI3-TH15-QR88",
			Name:  "Peach",
			Price: 2.99,
		},
		{
			Code:  "YRT6-72AS-K736-L4AR",
			Name:  "Green Pepper",
			Price: 0.79,
		},
		{
			Code:  "TQ4C-VV6T-75ZX-1RMR",
			Name:  "Gala Apple",
			Price: 3.59,
		},
	}
	return items
}
