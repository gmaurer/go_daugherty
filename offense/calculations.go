package offense

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"sync"
)

func ReadCSV() [][]string {
	file, err := os.Open("/Users/jas0126/dbs_go_lab/go_daugherty/go-stats-lab/offense/Batting.csv")

	if err != nil {
		log.Fatalf("Couldn't open file, error: %v", err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//The backend array isn't impacted either way, so creating a new slice is pretty minimal to performance
	return records[1:]
}

func ReadCSVToChannel(playerChannel chan []string) {
	file, err := os.Open("/Users/jas0126/dbs_go_lab/go_daugherty/go-stats-lab/offense/Batting.csv")

	if err != nil {
		log.Fatalf("Couldn't open file, error: %v", err)
	}

	defer file.Close()
	csvReader := csv.NewReader(file)

	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range records[1:] {
		playerChannel <- record
	}

}

func CreateBatterCollection(records [][]string) []Batter {
	battersList := make([]Batter, len(records))

	//Dynamically read header names to build struct?
	//encoding/csv -> read interpret the datatype
	for i, record := range records {
		batter := &battersList[i]
		batter.PlayerID = record[0]
		batter.YearID = convertStringToStat(record[1])
		batter.Stint = convertStringToStat(record[2])
		batter.TeamID = record[3]
		batter.LeagueID = record[4]
		batter.Games = convertStringToStat(record[5])
		batter.AtBats = convertStringToStat(record[6])
		batter.Runs = convertStringToStat(record[7])
		batter.Hits = convertStringToStat(record[8])
		batter.Doubles = convertStringToStat(record[9])
		batter.Triples = convertStringToStat(record[10])
		batter.Homeruns = convertStringToStat(record[11])
		batter.RunsBattedIn = convertStringToStat(record[12])
		batter.StolenBases = convertStringToStat(record[13])
		batter.CaughtStealing = convertStringToStat(record[14])
		batter.Walks = convertStringToStat(record[15])
		batter.Strikeouts = convertStringToStat(record[16])
		batter.IntentionalWalks = convertStringToStat(record[17])
		batter.HitByPitch = convertStringToStat(record[18])
		batter.SacrificeBunt = convertStringToStat(record[19])
		batter.SacrificeFly = convertStringToStat(record[20])
		batter.GroundIntoDoublePlay = convertStringToStat(record[21])
		//compare passing one pointer vs each
		batter.BattingAverage = calculateBattingAverage(batter)
		batter.OnBasePercentage = calculateOnBasePercentage(batter)
		batter.SluggingPercentage = calculateSluggingPercentage(batter)
		batter.OnBasePlusSlugging = calculateOnBasePlusSlugging(batter)
	}

	return battersList

}

func ProcessBatter(record []string) Batter {
	batter := Batter{
		record[0],
		convertStringToStat(record[1]),
		convertStringToStat(record[2]),
		record[3],
		record[4],
		convertStringToStat(record[5]),
		convertStringToStat(record[6]),
		convertStringToStat(record[7]),
		convertStringToStat(record[8]),
		convertStringToStat(record[9]),
		convertStringToStat(record[10]),
		convertStringToStat(record[11]),
		convertStringToStat(record[12]),
		convertStringToStat(record[13]),
		convertStringToStat(record[14]),
		convertStringToStat(record[15]),
		convertStringToStat(record[16]),
		convertStringToStat(record[17]),
		convertStringToStat(record[18]),
		convertStringToStat(record[19]),
		convertStringToStat(record[20]),
		convertStringToStat(record[21]),
		0,
		0,
		0,
		0,
	}
	batter.BattingAverage = calculateBattingAverage(&batter)
	batter.OnBasePercentage = calculateOnBasePercentage(&batter)
	batter.SluggingPercentage = calculateSluggingPercentage(&batter)
	batter.OnBasePlusSlugging = calculateOnBasePlusSlugging(&batter)

	return batter
}

func FanOutProcessing(records [][]string) []Batter {

	const numberOfGoroutines = 100
	segementSize := len(records) / numberOfGoroutines
	batterSlice := make([]Batter, len(records))

	var wg sync.WaitGroup

	for i := 0; i < numberOfGoroutines; i++ {
		startingIndex := i * segementSize
		endingIndex := (i + 1) * segementSize

		if i == numberOfGoroutines-1 {
			endingIndex = len(records)
		}

		wg.Add(1)

		go func(segment int) {
			for j, batter := range records[startingIndex:endingIndex] {
				batterSlice[segment*segementSize+j] = ProcessBatter(batter)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
	return batterSlice
}

func worker(playerChan <-chan []string, results chan<- Batter) {
	for record := range playerChan {
		results <- ProcessBatter(record)
	}
}

func WorkerPoolProcessing(records [][]string) []Batter {
	numJobs := len(records)
	playerChannel := make(chan []string, numJobs)
	results := make(chan Batter, numJobs)
	var playerSlice = make([]Batter, numJobs)

	numberOfWorkers := 32

	for i := 1; i <= numberOfWorkers; i++ {
		go worker(playerChannel, results)
	}

	go func() {
		for _, record := range records {
			playerChannel <- record
		}
	}()

	go func() {
		for j := 0; j < numJobs; j++ {
			playerSlice[j] = <-results
		}
	}()

	return playerSlice

}

func convertStringToStat(input string) int {
	if i, err := strconv.Atoi(input); err == nil {
		return i
	} else {
		return 0
	}
}

func GetAllBatterStats() {
	batters := CreateBatterCollection(ReadCSV())

	for _, batter := range batters {
		fmt.Println(batter)
	}
}

func GetAllBattersConcurrently() {
	batters := FanOutProcessing(ReadCSV())

	for _, batter := range batters {
		fmt.Println(batter)
	}
}

func GetAllBattersPool() {
	batters := WorkerPoolProcessing(ReadCSV())

	for _, batter := range batters {
		fmt.Println(batter)
	}
}

// func GetMVPPlayers() {
// 	batters := []Batter{
// 		{"troutmi01", 2014, 1, "LAA", "AL", 157, 602, 115, 173, 39, 9, 36, 111, 16, 2, 83, 184, 6, 10, 0, 10, 6, 0, 0, 0, 0},
// 		{"cabremi01", 2013, 1, "DET", "AL", 148, 555, 103, 193, 26, 1, 44, 137, 3, 0, 90, 94, 19, 5, 0, 2, 19, 0, 0, 0, 0},
// 		{"hamiljo03", 2010, 1, "TEX", "AL", 133, 518, 95, 186, 40, 3, 32, 100, 8, 1, 43, 95, 5, 5, 1, 4, 11, 0, 0, 0, 0},
// 		{"mauerjo01", 2009, 1, "MIN", "AL", 138, 523, 94, 191, 30, 1, 28, 96, 4, 1, 76, 63, 14, 2, 0, 5, 13, 0, 0, 0, 0},
// 	}

// 	for _, batter := range batters {
// 		calculateTotalStats(&batter)
// 		fmt.Println(batter)
// 	}

// }

// func calculateTotalStats(batterPtr *Batter) {
// 	calculateBattingAverage(batterPtr)
// 	calculateOnBasePercentage(batterPtr)
// 	calculateSluggingPercentage(batterPtr)
// 	calculateOnBasePlusSlugging(batterPtr)
// }

func calculateBattingAverage(batterPtr *Batter) float64 {
	if batterPtr.AtBats != 0 {
		return toFixed(float64(batterPtr.Hits)/float64(batterPtr.AtBats), 3)
	}
	return 0
}

func calculateOnBasePercentage(batterPtr *Batter) float64 {
	numer := float64(batterPtr.Hits + batterPtr.Walks + batterPtr.IntentionalWalks + batterPtr.HitByPitch)
	denom := float64(batterPtr.AtBats + batterPtr.Walks + batterPtr.IntentionalWalks + batterPtr.HitByPitch + batterPtr.SacrificeFly)

	if denom != 0 {
		return toFixed(numer/denom, 3)
	}
	return 0

}

func calculateSluggingPercentage(batterPtr *Batter) float64 {
	if batterPtr.AtBats != 0 {
		numer := float64(batterPtr.Hits + batterPtr.Doubles + (2 * batterPtr.Triples) + (3 * batterPtr.Homeruns))
		return toFixed(numer/float64(batterPtr.AtBats), 3)
	}
	return 0

}

func calculateOnBasePlusSlugging(batterPtr *Batter) float64 {
	return toFixed(batterPtr.OnBasePercentage+batterPtr.SluggingPercentage, 3)
}

func toFixed(stat float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(int(stat*output)) / output
}
