package main

import (
	"fmt"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func SanitizeBookTitle(title string) string {
	return cases.Title(language.Und).String(title)
}

func GiveFavorite(favorite string) string {
	return "My favorite book is " + favorite
}

func main() {
	fmt.Println(GiveFavorite("Frankenstein"))
}
