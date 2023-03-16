package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type basicplayer struct {
	PlayerID   string `db:"playerid"`
	BirthYear  int    `db:"birthyear"`
	BirthMonth int    `db:"birthmonth"`
	FirstName  string `db:"namefirst"`
	LastName   string `db:"namelast"`
}

type nullableplayer struct {
	PlayerID   string         `db:"playerid"`
	BirthYear  sql.NullInt64  `db:"birthyear"`
	BirthMonth sql.NullInt64  `db:"birthmonth"`
	FirstName  sql.NullString `db:"namefirst"`
	LastName   string         `db:"namelast"`
}

type basicplayerjoin struct {
	PlayerID     string `db:"playerid"`
	BirthYear    int    `db:"birthyear"`
	BirthMonth   int    `db:"birthmonth"`
	FirstName    string `db:"namefirst"`
	LastName     string `db:"namelast"`
	BattingStats []basicbatting
}

type basicbatting struct {
	PlayerID string `db:"playerid"`
	YearID   int    `db:"yearid"`
	TeamID   string `db:"teamid"`
	LeagueID string `db:"leagueid"`
	G        int    `db:"games"`
	AB       int    `db:"atbats"`
	H        int    `db:"hits"`
}

type player struct {
	PlayerID     string `db:"playerid"`
	BirthYear    int    `db:"birthyear"`
	BirthMonth   int    `db:"birthmonth"`
	BirthDay     int    `db:"birthday"`
	BirthCountry string `db:"birthcountry"`
	BirthState   string `db:"birthstate"`
	BirthCity    string `db:"birthcity"`
	DeathYear    int    `db:"deathyear"`
	DeathMonth   int    `db:"deathmonth"`
	DeathDay     int    `db:"deathday"`
	DeathCountry string `db:"deathcountry"`
	DeathState   string `db:"deathstate"`
	DeathCity    string `db:"deathcity"`
	FirstName    string `db:"namefirst"`
	LastName     string `db:"namelast"`
	GivenName    string `db:"namegiven"`
	Weight       int    `db:"weight"`
	Height       int    `db:"height"`
	Bat          string `db:"bat"`
	Throws       string `db:"throws"`
	Debut        string `db:"debut"`
	FinalGame    string `db:"finalgame"`
	RetroID      string `db:"retroid"`
	BbrefID      string `db:"bbrefid"`
}

func findPlayerById(id string) basicplayer {
	var p basicplayer
	if err := db.QueryRow("SELECT playerid, birthyear, birthmonth, namefirst, namelast FROM players WHERE playerid = $1",
		id).Scan(&p.PlayerID, &p.BirthYear, &p.BirthMonth, &p.FirstName, &p.LastName); err != nil {
		log.Fatalf("No player found with playerID of: %s", err)
	}

	return p

}

func findBasicPlayersByBirthYear(year int) []basicplayer {
	rows, err := db.Query("SELECT playerid, birthyear, birthmonth, namefirst, namelast FROM players WHERE birthyear = $1", year)

	if err != nil {
		log.Fatalf("Something went wrong with our query: %v", err)
	}

	var playerList []basicplayer

	for rows.Next() {
		var p basicplayer
		if err := rows.Scan(&p.PlayerID, &p.BirthYear, &p.BirthMonth, &p.FirstName, &p.LastName); err != nil {
			log.Fatal(err)
		}
		playerList = append(playerList, p)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	return playerList
}

func findBasicPlayersByBirthYearAndMonth(year int, month int) []basicplayer {
	rows, err := db.Query("SELECT playerid, birthyear, birthmonth, namefirst, namelast FROM players WHERE birthyear = $1 and birthmonth = $2", year, month)

	if err != nil {
		log.Fatal("Blah")
	}

	var playerList []basicplayer

	for rows.Next() {
		var p basicplayer
		if err := rows.Scan(&p.PlayerID, &p.BirthYear, &p.BirthMonth, &p.FirstName, &p.LastName); err != nil {
			log.Fatal(err)
		}
		playerList = append(playerList, p)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	return playerList
}

func createAPlayerExample() {
	_, err := db.Exec(`insert into "players" values('alonspe01', 1994, 12, 7, 'USA', 'FL', 'Tampa', 
	null, null, null, null, null, null, 'Peter', 'Alonso', 'Pete', 245, 75, 'R', 'R', '2019-03-28', null, 'alonspe01', 'alonspe01')`)

	if err != nil {
		log.Fatalf("Something went wrong while saving %v", err)
	}
}

func updateAPlayerExamle(givenName string, playerId string) {
	_, err := db.Exec(`update players set namegiven = $1 where playerid = $2`, givenName, playerId)
	if err != nil {
		log.Fatalf("Something went wrong while updating %v", err)
	}
}

func deletePlayerById(playerId string) {
	_, err := db.Exec(`delete from players where playerid = $1`, playerId)
	if err != nil {
		log.Fatalf("Something went wrong while updating %v", err)
	}
}

func findPlayerByIdJoin(id string) basicplayerjoin {
	var p basicplayerjoin
	if err := db.QueryRow("SELECT playerid, birthyear, birthmonth, namefirst, namelast FROM players WHERE playerid = $1",
		id).Scan(&p.PlayerID, &p.BirthYear, &p.BirthMonth, &p.FirstName, &p.LastName); err != nil {
		log.Fatalf("No player found with playerID of: %s", err)
	} else {

		rows, err := db.Query("SELECT playerid, yearid, teamid, leagueid, games, atbats, hits FROM batting WHERE playerid = $1", id)

		if err != nil {
			log.Fatal("Blah")
		}

		var battingstats []basicbatting

		for rows.Next() {
			var b basicbatting
			if err := rows.Scan(&b.PlayerID, &b.YearID, &b.TeamID, &b.LeagueID, &b.G, &b.AB, &b.H); err != nil {
				log.Fatal(err)
			}
			battingstats = append(battingstats, b)
		}
		p.BattingStats = battingstats
	}

	return p
}

func findPlayerByIdJoinAlt(id string) basicplayerjoin {

	rows, err := db.Query("SELECT p.playerid, p.birthyear, p.birthmonth, p.namefirst, p.namelast, b.playerid, b.yearid, b.teamid, b.leagueid, b.games, b.atbats, b.hits FROM players p JOIN batting b ON p.playerid = b.playerid WHERE p.playerid = $1 LIMIT 1", id)

	if err != nil {
		log.Fatal("Blah")
	}

	p := &basicplayerjoin{}

	for rows.Next() {
		b := basicbatting{}

		err := rows.Scan(&p.PlayerID, &p.BirthYear, &p.BirthMonth, &p.FirstName, &p.LastName, &b.PlayerID, &b.YearID, &b.TeamID, &b.LeagueID, &b.G, &b.AB, &b.H)

		if err != nil {
			log.Fatalf("Something went wrong %v", err)
		}
		p.BattingStats = append(p.BattingStats, b)
	}

	return *p

}

func findPlayerByIdUsingStatement(playerID string) basicplayer {
	// We usually would have our prepared statement saved elsewhere to be available across the app
	stmt, err := db.Prepare("SELECT playerid, birthyear, birthmonth, namefirst, namelast FROM players WHERE playerid = $1")
	if err != nil {
		log.Fatal(err)
	}

	var p basicplayer

	err = stmt.QueryRow(playerID).Scan(&p.PlayerID, &p.BirthYear, &p.BirthMonth, &p.FirstName, &p.LastName)
	if err != nil {
		//We will check for now rows, will discuss under handling errors
		if err == sql.ErrNoRows {
			log.Fatal("No rows returned")
		} else {
			log.Fatalf("Something went wrong: %v", err)
		}
	}
	return p
}

func updatePlayersInsideTx() (bool, error) {
	tx, err := db.Begin()

	if err != nil {
		log.Fatalf("Something went wrong with our transaction: %v", err)
	}

	_, err = tx.Exec(`update players set birthyear = 1995 where playerid='seageco01'`)

	if err != nil {
		tx.Rollback()
		return false, err
	}

	_, err = tx.Exec(`update players set birthyear = 1995 where playerid='castrmi01'`)

	// if something goes wrong, let's roll it back
	if err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	return true, nil

}

func findNullablePlayerById(id string) basicplayer {

	p := nullableplayer{}

	err := db.QueryRow("SELECT playerid, birthyear, birthmonth, namefirst, namelast FROM players WHERE playerid = $1",
		id).Scan(&p.PlayerID, &p.BirthYear, &p.BirthMonth, &p.FirstName, &p.LastName)

	bp := basicplayer{PlayerID: p.PlayerID, LastName: p.LastName}

	if err != nil {
		log.Fatalf("No player found with playerID of: %s", err)
	}

	if p.BirthYear.Valid {
		bp.BirthYear = int(p.BirthYear.Int64)
	} else {
		bp.BirthYear = 1900
	}

	if p.BirthMonth.Valid {
		bp.BirthMonth = int(p.BirthMonth.Int64)
	} else {
		bp.BirthMonth = 01
	}

	if p.FirstName.Valid {
		bp.FirstName = p.FirstName.String
	} else {
		bp.FirstName = "N/A"
	}

	return bp

}
