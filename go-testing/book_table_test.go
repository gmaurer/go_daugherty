package main

import "testing"

func TestSanitizeBookTitleTable(t *testing.T) {

	var cases = []struct {
		name     string
		input    string
		expected string
	}{
		{"Single word title", "Dracula", "Dracula"},
		{"Multiple word title", "The Great Gatsby", "The Great Gatsby"},
		{"Improper capitalization", "AniMaL fARm", "Animal Farm"},
		{"No capitalization", "the giving tree", "The Giving Tree"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			result := SanitizeBookTitle(tc.input)
			if result != tc.expected {
				t.Errorf("Return %s, but wanted %s", result, tc.expected)
			}
		})
	}
}

func TestSanitizeBookTitleTableParallel(t *testing.T) {
	t.Parallel()
	var cases = []struct {
		name     string
		input    string
		expected string
	}{
		{"Single word title", "Dracula", "Dracula"},
		{"Multiple word title", "The Great Gatsby", "The Great Gatsby"},
		{"Improper capitalization", "AniMaL fARm", "Animal Farm"},
		{"No capitalization", "the giving tree", "The Giving Tree"},
	}

	for _, tc := range cases {
		tc := tc // => prevents concurrency bug
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			result := SanitizeBookTitle(tc.input)
			if result != tc.expected {
				t.Errorf("Return %s, but wanted %s", result, tc.expected)
			}
		})
	}
}
