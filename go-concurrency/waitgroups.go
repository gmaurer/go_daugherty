package main

import (
	"fmt"
	"sync"
	"time"
)

func process(processNum int) {
	fmt.Printf("Now processing process #%d\n", processNum)
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("Completed processing of process #%d\n", processNum)
}

func ProcessManager() {
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		fmt.Println(i)
		// SEE NOTES AS TO WHY DO WE REASSIGN i TO ANOTHER VALUE j
		// comment out the reassignment to see the issue
		// i := i

		// go func() {
		// 	defer wg.Done()
		// 	process(i)
		// }()

		//We could solve this by passing the argument to the closure
		go func(val int) {
			defer wg.Done()
			process(val)
		}(i)
	}
	wg.Wait()
}
