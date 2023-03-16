package main

import (
	"log"
)

func HighestLayer() {
	highLayer() // <-- Recovers from panic
	log.Println("This will print")
}

func highLayer() {
	defer func() {
		val := recover()
		if val != nil {
			log.Println(val)
		}
	}()
	lowLayer() // <-- Calls function that panics
	log.Println("This won't print")
}

func lowLayer() {
	panics()
	log.Println("This won't print")
}

func panics() {
	panic("Something happened")
}
