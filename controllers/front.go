package controllers

import (
	"database/sql"
	"net/http"
)

func RegisterControllers(db *sql.DB) {
	oc := newOffenseController(db)
	pc := newPlayerController(db)

	http.Handle("/offense/", *oc)
	http.Handle("/players/", *pc)
}
