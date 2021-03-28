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

func TestGetItem(t *testing.T) {
	initDb()
	// Test valid request
	expected := getInitialTestItems()[0]
	actual, err := getItem("A12T-4GH7-QPL9-3N4M")
	if err != nil {
		t.Fatalf(`Unexpected error: %s`, err.Error())
	}
	if !expected.equals(actual) {
		t.Fatalf(`Expected: %+v Actual: %+v`, expected, actual)
	}

	// Test nonexistant produce code
	_, err = getItem("A12T-4GH7-QPL9-3N4Z")
	switch err.(type) {
	case nil:
		t.Fatalf(`Expected: ItemNotFoundError`)
	case *ItemNotFoundError:
	default:
		t.Fatalf(`Expected: ItemNotFoundError`)
	}
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
	// Test valid request 
	errors := addItems(list)
	if len(errors) > 0 {
		t.Fatalf(`Unexpected errors: %+v`, errors)
	}
	for _, expected := range list.Items {
		actual, exists := itemsByCode[expected.Code]
		if !exists || !expected.equals(actual) {
			t.Fatalf(`Expected: %+v Actual: %+v`, expected, actual)
		}
	}

	//Test inserting duplicate produce codes
	errors = addItems(list)
	for _, err := range errors {
		switch err.(type) {
		case nil:
			t.Fatalf(`Expected: ItemConflictError`)
		case *ItemConflictError:
		default:
			t.Fatalf(`Expected: ItemConflictError`)
		}
	}
}

func TestDeleteItem(t *testing.T) {
	initDb()
	// Test valid request
	err := deleteItem("A12T-4GH7-QPL9-3N4M")
	if err != nil {
		t.Fatalf(`Unexpected error: %s`, err.Error())
	}
	item, exists := itemsByCode["A12T-4GH7-QPL9-3N4M"]
	if exists {
		t.Fatalf(`Expected: nil Actual: %+v`, item)
	}

	// Test deleting non-existant produce code 
	err = deleteItem("XXXX-XXXX-XXXX-XXXX")
	switch err.(type) {
	case nil:
		t.Fatalf(`Expected: ItemNotFoundError`)
	case *ItemNotFoundError:
	default:
		t.Fatalf(`Expected: ItemNotFoundError`)
	}

	// Test deleting invalid produce code
	err = deleteItem("XXXX")
	switch err.(type) {
	case nil:
		t.Fatalf(`Expected: BadCodeError`)
	case *BadCodeError:
	default:
		t.Fatalf(`Expected: BadCodeError`)
	}
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

// Helper for checking item lists have the same values (order doesn't matter)
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

// Helper for checking multiple produce codes
func checkCodes(t *testing.T, codes []string, expected bool) {
	for _, code := range codes {
		if isValid := isValidCode(code); isValid != expected {
			t.Fatalf(`%s -> Expected: %t Actual: %t`, code, expected, isValid)
		}
	}
}

// Get initial produce items
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
