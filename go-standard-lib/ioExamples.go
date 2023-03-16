package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadAllExample() {

	// NewReader from the strings pkg returns a new Reader reading from the parameter string
	r := strings.NewReader("Here we are going to read a string using the io ReadAll method")
	b, err := io.ReadAll(r)
	if err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}

	// This will return a slice of bytes
	fmt.Print(b)

	// Let's use io.WriteString to print to stdout
	if _, err := io.WriteString(os.Stdout, string(b)); err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
}

func CopyAndPipeExample() {

	// here we create the reader and writer that are not connected via the pipe
	r, w := io.Pipe()

	// we use an anonymous function to write our string into the writer and close it
	go func() {
		fmt.Fprint(w, "io.Reader stream\n")
		w.Close()
	}()

	// since we used a pipe, we don't need to read the string into our Reader
	// and now we can use the copy to write it to the stdout
	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatalf("Something went wrong %v", err)
	}
}
