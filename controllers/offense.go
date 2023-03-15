package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
)

type offenseController struct {
	playerIDPattern *regexp.Regexp
	db              *sql.DB
}

func (oc offenseController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch r.URL.Path {
	case "/offense/hits":
		fmt.Println("hits")
	case "/offense/doubles":
		fmt.Println("doubles")
	default:
		fmt.Println("Default")
	}
}

func newOffenseController(db *sql.DB) *offenseController {
	return &offenseController{
		playerIDPattern: regexp.MustCompile(`^/offense/(\d+)/?`),
		db:              db,
	}
}
