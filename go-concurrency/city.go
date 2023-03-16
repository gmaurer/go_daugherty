package main

import "time"

type city struct {
	Name       string  `json:"name"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Population int64   `json:"population,string"`
	Rank       int64   `json:"rank,string"`
	State      string  `json:"state"`
	processed  bool
}

func processCity(record *city) {
	//simulate some complex process that takes some amount of time
	time.Sleep(10 * time.Millisecond)
	record.processed = true
}
