package main

import (
	"fmt"
	"math"
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

// Invalid produce code
type BadCodeError struct {
	Code string
	Err  error
}

// Produce code doesn't exist
type ItemNotFoundError struct {
	Code string
	Err  error
}

// Duplicate produce code
type ItemConflictError struct {
	Code string
	Err  error
}

// Implement chi.Binder interface to decode request payloads into struct
// Add custom validation logic to decoding process
func (item *Item) Bind(r *http.Request) error {
	if item.Name == "" {
		return fmt.Errorf("name is a required field")
	}
	if item.Code == "" {
		return fmt.Errorf("code is a required field")
	}
	// Assume items are not free
	if item.Price == 0 {
		return fmt.Errorf("price is a required field")
	}
	if !isValidCode(item.Code) {
		return fmt.Errorf("code is invalid")
	}
	if !isValidPrice(item.Price) {
		return fmt.Errorf("price can have at most 2 decimal places")
	}
	return nil
}

// Implement chi.Renderer interface to encode struct to response payload
func (*Item) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Check if items have equal values
func (item *Item) equals(other *Item) bool {
	return (item.Name == other.Name &&
		item.Code == other.Code &&
		item.Price == other.Price)
}

// Implement chi.Renderer interface to encode struct to response payload
func (*ItemList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// Implement chi.Binder interface to decode request payloads into struct
// Add custom validation logic to decoding process
func (list *ItemList) Bind(r *http.Request) error {
	if len(list.Items) < 1 {
		return fmt.Errorf("item list cannot be empty")
	}
	for _, item := range list.Items {
		err := item.Bind(r)
		if err != nil {
			return err
		}
	}
	return nil
}

// Check if produce code is valid
// Code must consist of 4 groups of 4 alphanumeric characters separated by dashes
// Codes are case insensitive
// e.g. ABC4-HDb8-8JSJ-KSDj
func isValidCode(code string) bool {
	codePattern := "^(?:[A-Za-z0-9]{4}-){3}[A-Za-z0-9]{4}$"
	matched, _ := regexp.MatchString(codePattern, code)
	return matched
}

// Checks that the price float has the valid precision
func isValidPrice(price float64) bool {
	// Arbitrary threshold for float comparison
	threshold := 1e-9
	// Price with digits after "cents place" removed
	expectedPrecisionPrice := math.Floor(price*100) / 100
	return math.Abs(price-expectedPrecisionPrice) < threshold
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
