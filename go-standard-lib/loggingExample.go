package main

import (
	"log"
)

func SimpleLogging() {
	log.Println("Here is a simple log")
}
func LoggingWithFormatSpecifiers(result any) {
	switch t := result.(type) {
	case int64:
		log.Printf("Log msg: %d", t)
	case float64:
		log.Printf("Log msg: %f", t)
	case string:
		log.Printf("Log msg: %s", t)
	default:
		log.Printf("Log msg: %v", t)
	}
}

func CountValidator(count int) {
	if count > 100 {
		log.Fatalf("Count exceeded upper limit with a value of %d, exiting", count)
	}
	log.Printf("Count was %d", count)
}

func CountPanic(count int) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()

	if count > 100 {
		log.Panicf("Count exceeded upper limit with a value of %d, exiting", count)
	}
}
