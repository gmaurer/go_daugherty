package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func readJSONDataToChan(cityChan chan city) {

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

	for _, city := range cities {
		cityChan <- city
	}
}

func channelWorker(citiesChan <-chan city, results chan<- city) {
	for x := range citiesChan {
		processCity(&x)
		results <- x
	}
}

func ChannelWorkerPoolExample() {
	const numJobs = 1000
	citiesChan := make(chan city, numJobs)
	results := make(chan city, numJobs)
	outputSlice := make([]city, numJobs)
	go readJSONDataToChan(citiesChan)

	workerPoolSize := 32

	for i := 1; i <= workerPoolSize; i++ {
		go channelWorker(citiesChan, results)
	}

	// this is only for printing purposes
	cityCounter := 0
	for a := 0; a < numJobs; a++ {
		cityCounter++
		outputSlice[a] = <-results
	}

}
