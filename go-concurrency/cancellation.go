package main

import "fmt"

func countDown(liftoff int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{}) // you could also make this a struct and close it rather than sending true
	cancel := func() {
		close(done)
	}

	go func() {
		for i := liftoff; i > 0; i-- {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
		//remember to close out
		close(ch)
	}()
	return ch, cancel
}

func CancellationExample() {
	ch, cancel := countDown(10)
	for i := range ch {
		if i < 1 {
			break
		}
		fmt.Printf("T-Minus %d seconds\n", i)
	}
	fmt.Println("Liftoff!")
	cancel()
}
