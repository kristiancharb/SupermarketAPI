package main

import "testing"

func TestGetHello(t *testing.T) {
	expected := "Hello"
	actual := GetHello()
	if expected != actual {
		t.Fatalf(`GetHello() = %q, want match for %#q`, actual, expected)
	}
}
