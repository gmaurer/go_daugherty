package main

import (
	"fmt"
	"log"
	"time"
)

// we can get the current time or set a dedicated time
var now time.Time = time.Now()
var newYearsDay time.Time = time.Date(2023, 1, 1, 1, 20, 30, 856161, time.Local)

func DisplayTimes() {
	p := fmt.Println

	fmt.Printf("Today's Date: %v\n", now)
	fmt.Printf("New Years Day: %v\n", newYearsDay)

	// getting various time components
	p(now.Year())
	p(now.Month())
	p(now.Day())
	p(now.Hour())
	p(now.Minute())
	p(now.Second())
	p(now.Nanosecond())
	p(now.Location())
	p(now.Weekday())

}

func TimeDurations() {

	p := fmt.Println

	// boolean methods that compare dates/times
	p(newYearsDay.Before(now))
	p(newYearsDay.After(now))
	p(newYearsDay.Equal(now))

	// sub gives the interval between two times.
	diff := now.Sub(newYearsDay)
	p(diff)

	// we can provide the time components of the interval duration
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Add allows us to add time duration to a given duration
	p(newYearsDay.Add(diff))

	// Using a negative duration moves backward from a given duration
	p(newYearsDay.Add(-diff))
}

func TimerExample() {
	timer := time.NewTimer(0 * time.Second)

	<-timer.C
	fmt.Println("Timer  fired")
	time.Sleep(5 * time.Second)
	timer.Stop()
	fmt.Println("Timer stopped")
}

func TickerExample() {

	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	// We use a channel to get each interval
	// we will talk channels in the concurrency section
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}

func FormatTime() {
	t, err := time.Parse(time.UnixDate, "Wed Feb 15 12:00:00 CST 2023")
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}

	fmt.Println("Time formatted:", t)
	fmt.Println("Time in UTC:", t.UTC().Format(time.UnixDate))
}
