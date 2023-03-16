package main

import (
	"fmt"
	"time"
)

func DoneChannelExample() {
	// every one second ticket.C writes to the channel
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)

	// here is our for-select-done
	// if the done channel is written to, it signals to return and exit the goroutine
	go func() {
		for {
			select {
			case <-done:
				fmt.Println("I've been signaled, my time has come")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
}
