package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	"source.cloud.google.com/dl-gcp-cngo-sbox-devenv-b1/go_daugherty/go-stats-lab/controllers"
)

type connection struct {
	Driver   string
	Host     string
	User     string
	Password string
	Dbname   string
}

var db *sql.DB

func init() {

	//update with your file path
	configFile, err := os.Open("/Users/gkvrg/Documents/projects/go_daugherty/go-db/dbconfig.json")

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer configFile.Close()

	configBytes, _ := io.ReadAll(configFile)
	conn := new(connection)
	err = json.Unmarshal(configBytes, &conn)

	if err != nil {
		log.Fatalf("Error creating config: %v", err)
	}

	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", conn.Host, conn.User, conn.Password, conn.Dbname)

	db, err = sql.Open(conn.Driver, connString)

	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Erroring when pinging the database: %v", err)
	}

}

func main() {

	// service.FindPlayerAndTotalDesiredStat("dunnad01", "so", db)
	controllers.RegisterControllers(db)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
