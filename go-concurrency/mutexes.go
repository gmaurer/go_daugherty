package main

import (
	"fmt"
	"sync"
)

var counter int = 0

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	counter++
	m.Unlock()
	wg.Done()
}

func MutexExample() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go increment(&wg, &m)
	}

	wg.Wait()
	fmt.Println("Value of counter:", counter)
}
