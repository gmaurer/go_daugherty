package main

import "fmt"

func FmtExamples() {

	name := "World"

	fmt.Print("Hello", name)
	fmt.Println("Hello", name)
	fmt.Printf("Hello %s\n", name)

	msg := fmt.Sprintf("Hello %s!", name)
	fmt.Println(msg)
}
