package main

import "testing"

func TestGetHello(t *testing.T) {
	expected := "Hello"
	actual := getHello()
	if expected != actual {
		t.Fatalf(`GetHello() = %q, want match for %#q`, actual, expected)
	}
}
