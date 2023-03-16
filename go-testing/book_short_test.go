package main

import (
	"testing"
)

func TestSanitizeBookTitle(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration in short mode")
	}

	result := SanitizeBookTitle("ThE GrEEn miLe")
	if result != "The Green Mile" {
		t.Errorf("Return %s, but wanted The Green Mile", result)
	}
}
