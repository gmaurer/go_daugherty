package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func AtomicCounter() {
	var counter uint64
	var waitGroup sync.WaitGroup

	for i := 0; i < 20; i++ {
		waitGroup.Add(1)

		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddUint64(&counter, 1)
			}
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
	fmt.Printf("Atomic Counter Value: %d\n", counter)
}

func AtomicCounterDuringUpdates() {
	var counter uint64
	var waitGroup sync.WaitGroup

	for i := 0; i < 20; i++ {
		waitGroup.Add(1)

		go func() {
			for j := 0; j < 100; j++ {
				atomic.AddUint64(&counter, 1)
			}
			fmt.Printf("Atomic Counter Mid Update Value: %d\n", atomic.LoadUint64(&counter))
			waitGroup.Done()
		}()
	}

	waitGroup.Wait()
}
