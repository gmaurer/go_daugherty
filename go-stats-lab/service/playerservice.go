package service

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"golang.org/x/exp/maps"
)

type NullSafePlayer struct {
	PlayerID     string  `db:"playerid"`
	BirthYear    *int    `db:"birthyear"`
	BirthMonth   *int    `db:"birthmonth"`
	BirthDay     *int    `db:"birthday"`
	BirthCountry *string `db:"birthcountry"`
	BirthState   *string `db:"birthstate"`
	BirthCity    *string `db:"birthcity"`
	DeathYear    *int    `db:"deathyear"`
	DeathMonth   *int    `db:"deathmonth"`
	DeathDay     *int    `db:"deathday"`
	DeathCountry *string `db:"deathcountry"`
	DeathState   *string `db:"deathstate"`
	DeathCity    *string `db:"deathcity"`
	FirstName    *string `db:"namefirst"`
	LastName     *string `db:"namelast"`
	GivenName    *string `db:"namegiven"`
	Weight       *int    `db:"weight"`
	Height       *int    `db:"height"`
	Bat          *string `db:"bat"`
	Throws       *string `db:"throws"`
	Debut        *string `db:"debut"`
	FinalGame    *string `db:"finalgame"`
	RetroID      *string `db:"retroid"`
	BbrefID      *string `db:"bbrefid"`
	Stats        []*Batting
}

type Player struct {
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
	Stats        []Batting
}

type PlayerSimple struct {
	PlayerID  string `db:"playerid"`
	FirstName string `db:"namefirst"`
	LastName  string `db:"namelast"`
	Batting   []BattingSimple
}

type PlayerAndStat struct {
	PlayerID  string `db:"playerid"`
	FirstName string `db:"namefirst"`
	LastName  string `db:"namelast"`
	Batting   []BattingTop
}

type BattingSimple struct {
	PlayerID string `db:"playerid"`
	YearID   int    `db:"yearid"`
	TeamID   string `db:"teamid"`
	Homeruns int    `db:"homeruns"`
}

type Batting struct {
	PlayerID             string `db:"playerid"`
	YearID               int    `db:"yearid"`
	Stint                int    `db:"stint"`
	TeamID               string `db:"teamid"`
	LeagueID             string `db:"leagueid"`
	Games                int    `db:"games"`
	AtBats               int    `db:"atbats"`
	Runs                 int    `db:"runs"`
	Hits                 int    `db:"hits"`
	Doubles              int    `db:"doubles"`
	Triples              int    `db:"triples"`
	Homeruns             int    `db:"homeruns"`
	RunsBattedIn         int    `db:"rbis"`
	StolenBases          int    `db:"stolenbases"`
	CaughtStealing       int    `db:"caughtstealing"`
	Walks                int    `db:"walks"`
	Strikeouts           int    `db:"strikeouts"`
	IntentionalWalks     int    `db:"intentionalwalks"`
	HitByPitch           int    `db:"hitbypitch"`
	SacrificeBunt        int    `db:"sacrificebunt"`
	SacrificeFly         int    `db:"sacrificeflyout"`
	GroundIntoDoublePlay int    `db:"groundintodoubleplay"`
	BattingAverage       float64
	OnBasePercentage     float64
	SluggingPercentage   float64
	OnBasePlusSlugging   float64
}

type BattingTop struct {
	PlayerID string `db:"playerid"`
	YearID   int    `db:"yearid"`
	Hits     int    `db:"hits"`
}

type PlayersError struct {
	StatusCode int
	ErrMsg     *string
}

type PlayerAndHomeruns struct {
	PlayerID    string
	FirstName   string
	LastName    string
	NumberOfHRs int
	Year        int
	Team        string
}

type Nullableplayer struct {
	PlayerID   string         `db:"playerid"`
	BirthYear  sql.NullInt64  `db:"birthyear"`
	BirthMonth sql.NullInt64  `db:"birthmonth"`
	FirstName  sql.NullString `db:"namefirst"`
	LastName   string         `db:"namelast"`
}

type PlayerAndTopStat struct {
	FirstName     string
	LastName      string
	RequestedStat string
	TopStat       int
}

func FindAllPlayers(db *sql.DB) ([]*NullSafePlayer, PlayersError) {

	rows, err := db.Query(`SELECT * FROM players`)

	if err != nil {
		msg := err.Error()
		return nil, PlayersError{500, &msg}
	}

	var playerList []*NullSafePlayer

	for rows.Next() {
		p := &NullSafePlayer{}

		err := rows.Scan(&p.PlayerID, &p.BirthYear, &p.BirthMonth, &p.BirthDay,
			&p.BirthCountry, &p.BirthState, &p.BirthCity, &p.DeathYear,
			&p.DeathMonth, &p.DeathDay, &p.DeathCountry, &p.DeathState,
			&p.DeathCity, &p.FirstName, &p.LastName, &p.GivenName, &p.Weight,
			&p.Height, &p.Bat, &p.Throws, &p.Debut, &p.FinalGame, &p.RetroID, &p.BbrefID,
		)

		if err != nil {
			msg := err.Error()
			return nil, PlayersError{500, &msg}
		}
		playerList = append(playerList, p)
	}
	defer rows.Close()

	return playerList, PlayersError{200, nil}
}

func FindAllPlayersAlt(db *sql.DB) ([]*Player, PlayersError) {

	queryString := `SELECT
		playerid,
		COALESCE(birthyear, 0),
		COALESCE(birthmonth, 0),
		COALESCE(birthday, 0),
		COALESCE(birthcountry, 'N/A'),
		COALESCE(birthstate, 'N/A'),
		COALESCE(birthcity, 'N/A'),
		COALESCE(deathyear, 0),
		COALESCE(deathmonth, 0),
		COALESCE(deathday, 0),
		COALESCE(deathcountry, 'N/A'),
		COALESCE(deathstate, 'N/A'),
		COALESCE(deathcity, 'N/A'),
		COALESCE(namefirst, 'N/A'),
		COALESCE(namelast, 'N/A'),
		COALESCE(namegiven, 'N/A'),
		COALESCE(weight, 0),
		COALESCE(height, 0),
		COALESCE(bat, 'U'),
		COALESCE(throws, 'U'),
		COALESCE(debut, TO_DATE('1800/01/01', 'YYYY/MM/DD')),
		COALESCE(finalgame, TO_DATE('1800/01/01', 'YYYY/MM/DD')),
		COALESCE(retroid, 'N/A'),
		COALESCE(bbrefid, 'N/A')
	FROM players`

	rows, err := db.Query(queryString)

	if err != nil {
		msg := err.Error()
		return nil, PlayersError{500, &msg}
	}

	var playerList []*Player

	for rows.Next() {
		p := &Player{}

		err := rows.Scan(&p.PlayerID, &p.BirthYear, &p.BirthMonth, &p.BirthDay,
			&p.BirthCountry, &p.BirthState, &p.BirthCity, &p.DeathYear,
			&p.DeathMonth, &p.DeathDay, &p.DeathCountry, &p.DeathState,
			&p.DeathCity, &p.FirstName, &p.LastName, &p.GivenName, &p.Weight,
			&p.Height, &p.Bat, &p.Throws, &p.Debut, &p.FinalGame, &p.RetroID, &p.BbrefID,
		)

		if err != nil {
			msg := err.Error()
			return nil, PlayersError{500, &msg}
		}
		playerList = append(playerList, p)
	}

	defer rows.Close()

	return playerList, PlayersError{200, nil}
}

func FindAllPlayersByTeamAndSeasonYear(db *sql.DB, team string, year int) ([]*NullSafePlayer, PlayersError) {

	rows, err := db.Query(`SELECT * FROM players p INNER JOIN batting b ON p.playerid = b.playerid where b.teamid = $1 and b.yearid = $2`,
		team, year)

	if err != nil {
		msg := err.Error()
		return nil, PlayersError{500, &msg}
	}

	defer rows.Close()

	playerMap := make(map[string]*NullSafePlayer)

	for rows.Next() {
		p := &NullSafePlayer{}
		b := &Batting{}
		err := rows.Scan(&p.PlayerID, &p.BirthYear, &p.BirthMonth, &p.BirthDay,
			&p.BirthCountry, &p.BirthState, &p.BirthCity, &p.DeathYear,
			&p.DeathMonth, &p.DeathDay, &p.DeathCountry, &p.DeathState,
			&p.DeathCity, &p.FirstName, &p.LastName, &p.GivenName, &p.Weight,
			&p.Height, &p.Bat, &p.Throws, &p.Debut, &p.FinalGame, &p.RetroID, &p.BbrefID,
			&b.PlayerID, &b.YearID, &b.Stint, &b.TeamID, &b.LeagueID, &b.Games, &b.AtBats, &b.Runs, &b.Hits,
			&b.Doubles, &b.Triples, &b.Homeruns, &b.RunsBattedIn, &b.StolenBases, &b.CaughtStealing, &b.Walks, &b.Strikeouts,
			&b.IntentionalWalks, &b.HitByPitch, &b.SacrificeBunt, &b.SacrificeFly, &b.GroundIntoDoublePlay)

		if err != nil {
			msg := err.Error()
			return nil, PlayersError{500, &msg}
		}

		_, ok := playerMap[p.PlayerID]

		if ok {
			playerMap[p.PlayerID].Stats = append(playerMap[p.PlayerID].Stats, b)
		} else {
			p.Stats = append(p.Stats, b)
			playerMap[p.PlayerID] = p
		}
	}

	if len(maps.Values(playerMap)) == 0 {
		msg := "No Team Found"
		return nil, PlayersError{404, &msg}
	}
	return maps.Values(playerMap), PlayersError{200, nil}
}

func FindAllPlayersWithStats(db *sql.DB) ([]*NullSafePlayer, PlayersError) {
	rows, err := db.Query(`SELECT * FROM players p INNER JOIN batting b ON p.playerid = b.playerid`)

	if err != nil {
		msg := "Something went wrong"
		return nil, PlayersError{500, &msg}
	}

	defer rows.Close()

	playerMap := make(map[string]*NullSafePlayer)

	for rows.Next() {
		p := &NullSafePlayer{}
		b := &Batting{}
		err := rows.Scan(&p.PlayerID, &p.BirthYear, &p.BirthMonth, &p.BirthDay,
			&p.BirthCountry, &p.BirthState, &p.BirthCity, &p.DeathYear,
			&p.DeathMonth, &p.DeathDay, &p.DeathCountry, &p.DeathState,
			&p.DeathCity, &p.FirstName, &p.LastName, &p.GivenName, &p.Weight,
			&p.Height, &p.Bat, &p.Throws, &p.Debut, &p.FinalGame, &p.RetroID, &p.BbrefID,
			&b.PlayerID, &b.YearID, &b.Stint, &b.TeamID, &b.LeagueID, &b.Games, &b.AtBats, &b.Runs, &b.Hits,
			&b.Doubles, &b.Triples, &b.Homeruns, &b.RunsBattedIn, &b.StolenBases, &b.CaughtStealing, &b.Walks, &b.Strikeouts,
			&b.IntentionalWalks, &b.HitByPitch, &b.SacrificeBunt, &b.SacrificeFly, &b.GroundIntoDoublePlay)

		if err != nil {
			msg := "Something went wrong"
			return nil, PlayersError{500, &msg}
		}

		_, ok := playerMap[p.PlayerID]

		if ok {
			playerMap[p.PlayerID].Stats = append(playerMap[p.PlayerID].Stats, b)
		} else {
			p.Stats = append(p.Stats, b)
			playerMap[p.PlayerID] = p
		}
	}

	if len(maps.Values(playerMap)) == 0 {
		msg := "No Team Found"
		return nil, PlayersError{404, &msg}
	}
	return maps.Values(playerMap), PlayersError{200, nil}
}

func FindPlayerById(playerId string, db *sql.DB) (*NullSafePlayer, PlayersError) {

	rows, err := db.Query(`SELECT * FROM players p INNER JOIN batting b ON p.playerid = b.playerid where p.playerid = $1`, playerId)

	if err != nil {
		msg := err.Error()
		return nil, PlayersError{500, &msg}
	}

	defer rows.Close()

	playerMap := make(map[string]*NullSafePlayer)

	for rows.Next() {
		p := &NullSafePlayer{}
		b := &Batting{}
		err := rows.Scan(&p.PlayerID, &p.BirthYear, &p.BirthMonth, &p.BirthDay,
			&p.BirthCountry, &p.BirthState, &p.BirthCity, &p.DeathYear,
			&p.DeathMonth, &p.DeathDay, &p.DeathCountry, &p.DeathState,
			&p.DeathCity, &p.FirstName, &p.LastName, &p.GivenName, &p.Weight,
			&p.Height, &p.Bat, &p.Throws, &p.Debut, &p.FinalGame, &p.RetroID, &p.BbrefID,
			&b.PlayerID, &b.YearID, &b.Stint, &b.TeamID, &b.LeagueID, &b.Games, &b.AtBats, &b.Runs, &b.Hits,
			&b.Doubles, &b.Triples, &b.Homeruns, &b.RunsBattedIn, &b.StolenBases, &b.CaughtStealing, &b.Walks, &b.Strikeouts,
			&b.IntentionalWalks, &b.HitByPitch, &b.SacrificeBunt, &b.SacrificeFly, &b.GroundIntoDoublePlay)

		if err != nil {
			msg := err.Error()
			return nil, PlayersError{500, &msg}
		}

		_, ok := playerMap[p.PlayerID]

		if ok {
			playerMap[p.PlayerID].Stats = append(playerMap[p.PlayerID].Stats, b)
		} else {
			p.Stats = append(p.Stats, b)
			playerMap[p.PlayerID] = p
		}
	}

	if len(maps.Values(playerMap)) == 0 {
		msg := "Player Not Found"
		return nil, PlayersError{404, &msg}
	}
	return maps.Values(playerMap)[0], PlayersError{200, nil}
}

func FindPlayerAndTotalDesiredStat(playerID string, requestedStat string, db *sql.DB) (*PlayerAndTopStat, PlayersError) {

	//create a whitelist
	possibleStats := map[string]string{
		"G":    "b.games",
		"AB":   "b.atbats",
		"R":    "b.runs",
		"H":    "b.hits",
		"2B":   "b.doubles",
		"3B":   "b.triples",
		"HR":   "b.homeruns",
		"RBI":  "b.rbis",
		"SB":   "b.stolenbases",
		"CS":   "b.caughtstealing",
		"BB":   "b.walks",
		"SO":   "b.strikeouts",
		"IBB":  "b.intentionalwalks",
		"HBP":  "b.hitbyptich",
		"SH":   "b.sacrificebunt",
		"SF":   "b.sacrificeflyout",
		"GIDP": "b.groundedintodoubleplay",
	}

	wlStat, ok := possibleStats[strings.ToUpper(requestedStat)]

	if !ok {
		log.Fatal("Dennis Nedry saying 'ah ah ah, you don't say the magic word'")
	}

	sqlSub := fmt.Sprintf("SELECT p.namefirst, p.namelast, SUM(%s) as totalHits FROM players p JOIN batting b ON p.playerid = b.playerid WHERE p.playerid = $1 GROUP BY p.playerid, p.namefirst, p.namelast, b.playerid", wlStat)

	p := &PlayerAndTopStat{}
	p.RequestedStat = requestedStat
	err := db.QueryRow(sqlSub, playerID).Scan(&p.FirstName, &p.LastName, &p.TopStat)

	if err != nil {
		msg := err.Error()
		if err == sql.ErrNoRows {
			return nil, PlayersError{404, &msg}
		} else {
			return nil, PlayersError{500, &msg}
		}
	}

	return p, PlayersError{200, nil}
}

//Deprecated Methods, need to update

func FindAllPlayerStatsById(playerPtr *NullSafePlayer, db *sql.DB) {

	statrows, err := db.Query(`SELECT * FROM batting where playerid = $1 order by yearid asc`, playerPtr.PlayerID)

	if err != nil {
		log.Fatalf("Something went wrong with our stats query: %v", err)
	}

	var statsList []*Batting

	for statrows.Next() {
		b := &Batting{}
		if err := statrows.Scan(&b.PlayerID, &b.YearID, &b.Stint, &b.TeamID, &b.LeagueID, &b.Games, &b.AtBats, &b.Runs, &b.Hits,
			&b.Doubles, &b.Triples, &b.Homeruns, &b.RunsBattedIn, &b.StolenBases, &b.CaughtStealing, &b.Walks, &b.Strikeouts,
			&b.IntentionalWalks, &b.HitByPitch, &b.SacrificeBunt, &b.SacrificeFly, &b.GroundIntoDoublePlay); err != nil {
			log.Fatal(err)
		}

		statsList = append(statsList, b)
	}

	statrows.Close()

	playerPtr.Stats = statsList

	//Need to do calculations
	// BattingAverage   float64
	// OnBasePercentage     float64
	// SluggingPercentage   float64
	// OnBasePlusSlugging   float64

}

func FindPlayerAndHomerunsById(playerID string, db *sql.DB) (*PlayerAndHomeruns, PlayersError) {
	p := &PlayerAndHomeruns{}

	err := db.QueryRow("SELECT p.playerid, p.namefirst, p.namelast, b.homeruns, b.yearid, b.teamid FROM players p JOIN batting b ON p.playerid = b.playerid WHERE p.playerid = $1 ORDER BY b.homeruns desc LIMIT 1",
		playerID).Scan(&p.PlayerID, &p.FirstName, &p.LastName, &p.NumberOfHRs, &p.Year, &p.Team)

	var msg string

	if err != nil {
		msg = err.Error()
		if err == sql.ErrNoRows {
			return nil, PlayersError{404, &msg}
		} else {
			return nil, PlayersError{500, &msg}
		}
	}

	return p, PlayersError{200, nil}
}
