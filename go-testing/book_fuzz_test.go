package main

import "testing"

func FuzzSanitizeBookTitle(f *testing.F) {
	f.Add("Hello world")
	f.Fuzz(func(t *testing.T, name string) {
		SanitizeBookTitle(name)
	})
}
