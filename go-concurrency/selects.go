package main

import (
	"fmt"
	"log"
	"time"
)

func SelectExample() {

	start := time.Now()
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	c4 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- 2
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c3 <- 3
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c4 <- 4
	}()

	for i := 0; i < 4; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("%d received\n", msg1)
		case msg2 := <-c2:
			fmt.Printf("%d received\n", msg2)
		case msg3 := <-c3:
			fmt.Printf("%d received\n", msg3)
		case msg4 := <-c4:
			fmt.Printf("%d received\n", msg4)

		}
	}
	log.Printf("Exectution time : %v", time.Since(start))
}
