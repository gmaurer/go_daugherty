package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func SerialProcessCities() {
	cities := readJSONDataSerial()
	processedCities := make([]city, 1000)

	for i, city := range cities {
		processCity(&city)
		processedCities[i] = city
	}
}

func readJSONDataSerial() []city {

	//update your path
	cityContent, err := os.Open("/Users/jas0126/dbs_go_lab/go_daugherty/go-concurrency/2013Cities.json")

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	defer cityContent.Close()

	cityBytes, _ := io.ReadAll(cityContent)
	var cities []city
	err = json.Unmarshal(cityBytes, &cities)

	if err != nil {
		log.Fatalf("Error unmarshalling JSON file: %v", err)
	}

	return cities
}
