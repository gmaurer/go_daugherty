package main

// "errors"
// "log"
import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	// HighestLayer()
	//SimplePanicAndRecoverExample()
	//ReceivesError()
	a := "Hello"
	val, err := strconv.Atoi(a)

	if err != nil {
		log.Fatal("Encountered error, exiting")
		// fmt.Println("I have an error")
	}

	fmt.Println(val)

	// err := CheckCount(125)
	// if err != nil {
	// 	log.Printf("There was an error: %s", err.Error())
	// 	return
	// }

	// msg, err := GetResponse(900)
	// if err != nil {
	// 	// msg is probably not correctly initialized
	// 	// so we shouldn't use it
	// 	log.Printf("There was an error: %s", err.Error())
	// 	log.Fatalf("error in get response: %v", err)
	// }
	// // there was no error, do something interesting with msg
	// log.Println(msg)

	// err := WrappedCount(125)
	// if err != nil {
	// 	log.Printf("Error: %s", err.Error())
	// 	log.Printf("Unwrapped Error: %s", errors.Unwrap(err))
	// 	return
	// }
}
