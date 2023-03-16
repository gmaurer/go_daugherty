package main

import (
	"log"
	"os"
)

func WriteFileContents(name string, data []byte) {
	f, err := os.OpenFile(name,
		os.O_CREATE|os.O_WRONLY|os.O_APPEND,
		0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n, err := f.Write(data)
	if err != nil {
		log.Printf("There was an error writing to the file: %s", err.Error())
	} else {
		log.Printf("Successfully wrote %d bytes to %s", n, name)
	}
}

func ReadFileContents(name string) []byte {
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}

	return data
}
