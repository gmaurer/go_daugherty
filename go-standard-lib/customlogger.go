package main

import (
	"log"
	"os"
)

var CustomLogger *log.Logger

// Here I'm using a special function in Go called "init" that gets called on the initialization of main
func init() {
	CustomLogger = log.New(os.Stdout, ":-( Custom Logger:\t", log.Ldate|log.Ltime)
}
