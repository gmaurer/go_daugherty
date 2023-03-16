package main

import "testing"

func TestGiveFavorite(t *testing.T) {
	result := GiveFavorite("Frankenstein")
	expected := "My favorite book is Frankenstein"

	if result != expected {
		t.Errorf("Return %s, but wanted %s", result, expected)
	}
}
