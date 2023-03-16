package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"source.cloud.google.com/dl-gcp-cngo-sbox-devenv-b1/go_daugherty/go-stats-lab/service"
)

type PlayersResponse struct {
	Response interface{}
	ErrMsg   *string
}

type playerController struct {
	playerServicePattern *regexp.Regexp
	playerIDPattern      *regexp.Regexp
	teamSeasonPattern    *regexp.Regexp
	db                   *sql.DB
}

func (pc playerController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sub := pc.playerServicePattern.FindStringSubmatch(r.URL.Path)[1]

	switch sub {
	case "id":
		fmt.Println(pc.playerIDPattern.FindStringSubmatch(r.URL.Path)[2])
		id := pc.playerIDPattern.FindStringSubmatch(r.URL.Path)[1]
		best := pc.playerIDPattern.FindStringSubmatch(r.URL.Path)[2]
		if len(best) != 0 {
			pc.getPlayerAndTotalStatById(id, best, w, r)
		} else {
			pc.getPlayerById(id, w, r)
		}
	case "all":
		pc.getAllPlayers(w, r)
	case "team":
		team := pc.teamSeasonPattern.FindStringSubmatch(r.URL.Path)[1]
		year, err := strconv.Atoi(pc.teamSeasonPattern.FindStringSubmatch(r.URL.Path)[2])
		if err != nil {
			year = 0
		}
		pc.getAllPlayersByTeamAndSeason(team, year, w, r)
	default:
		msg := "Method Not Found"
		responseHandler(nil, service.PlayersError{StatusCode: 405, ErrMsg: &msg}, w)
	}
}

func (pc playerController) getAllPlayers(w http.ResponseWriter, r *http.Request) {
	res, err := service.FindAllPlayersWithStats(pc.db)
	responseHandler(res, err, w)
}

func (pc playerController) getAllPlayersByTeamAndSeason(team string, year int, w http.ResponseWriter, r *http.Request) {
	res, err := service.FindAllPlayersByTeamAndSeasonYear(pc.db, team, year)
	responseHandler(res, err, w)
}

func (pc playerController) getPlayerById(id string, w http.ResponseWriter, r *http.Request) {
	res, err := service.FindPlayerById(id, pc.db)
	responseHandler(res, err, w)
}

func (pc playerController) getPlayerAndTotalStatById(id string, desiredstat string, w http.ResponseWriter, r *http.Request) {
	res, err := service.FindPlayerAndTotalDesiredStat(id, desiredstat, pc.db)
	responseHandler(res, err, w)
}

func newPlayerController(db *sql.DB) *playerController {
	return &playerController{
		playerServicePattern: regexp.MustCompile(`^/players/([a-zA-Z0-9]+)/?`),
		playerIDPattern:      regexp.MustCompile(`^/players/id/([a-zA-Z0-9]+)/?([a-zA-Z0-9]+)?`),
		teamSeasonPattern:    regexp.MustCompile(`^/players/team/([a-zA-Z0-9]+)/year/([0-9]{4})/?`),
		db:                   db,
	}
}

func responseHandler(res interface{}, error service.PlayersError, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if error.StatusCode == 200 {
		w.WriteHeader(http.StatusOK)
	} else if error.StatusCode == 204 {
		w.WriteHeader(http.StatusNotFound)
	} else if error.StatusCode == 404 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	encodeResponseAsJSON(PlayersResponse{res, error.ErrMsg}, w)

}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
