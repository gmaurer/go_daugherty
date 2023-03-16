package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"
)

func FanOutExample() {
	cities := readJSONData()
	const numberOfGoroutines = 100
	segmentSize := len(cities) / numberOfGoroutines
	processedCities := make([]city, len(cities))
	var wg sync.WaitGroup

	for i := 0; i < numberOfGoroutines; i++ {
		startingIndex := i * segmentSize
		endingIndex := (i + 1) * segmentSize

		//set endingIndex on last Goroutine if not even
		if i == numberOfGoroutines-1 {
			endingIndex = len(cities)
		}

		wg.Add(1)

		// pass an argument into the closure to prevent concurrency bug
		go func(segment int) {
			for j, city := range cities[startingIndex:endingIndex] {
				processCity(&city)
				processedCities[segment*segmentSize+j] = city
			}
			wg.Done()
		}(i)
	}

	wg.Wait()

}

func readJSONData() []city {

	//update your path
	cityContent, err := os.Open("/Users/gkvrg/Documents/projects/go_daugherty/go-concurrency/2013Cities.json")

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
