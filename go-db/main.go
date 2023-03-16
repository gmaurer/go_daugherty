package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	_ "github.com/lib/pq"
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
	fmt.Println(conn)

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

	// player := findPlayerById("zobribe01")
	// fmt.Println(player)

	// players := findBasicPlayersByBirthYear(1988)
	// fmt.Println(players)

	// otherPlayers := findBasicPlayersByBirthYearAndMonth(1988, 1)
	// fmt.Println(otherPlayers)

	createAPlayerExample()

	// updateAPlayerExamle("Peter", "alonspe01")

	//deletePlayerById("alonspe01")

	// player := findPlayerByIdJoin("suzukic01")
	// fmt.Println(player)

	// player := findPlayerByIdJoinAlt("suzukic01")
	// fmt.Println(player)

	player := findNullablePlayerById("bolan01")
	fmt.Println(player)

	// player := findPlayerByIdJoinAlt("suzukic01")
	// fmt.Println(player)

	// player := findPlayerByIdUsingStatement("suzukic01")
	// fmt.Println(player)

	defer db.Close()
}
