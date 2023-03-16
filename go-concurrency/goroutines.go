package main

// func LaunchCountdown(launchTime int) <-chan int {
// 	liftOff := make(chan int)
// 	go func() {
// 		for launchTime > 0 {
// 			liftOff <- launchTime
// 			launchTime--
// 			time.Sleep(1 * time.Second)
// 		}
// 		close(liftOff)
// 	}()
// 	return liftOff
// }

func main() {
	//for i := range LaunchCountdown(10) {
	//	fmt.Println(i)
	//}
	// fmt.Println("example")
	// ChannelExample()
	// BufferedChannelExample()
	// RangeBufferedChannelExample()
	// IdiomaticBufferedChannelExample()
	//ProcessManager()
	// ChannelBlockingExample()
	// BufferedChannelBlockingExample()
	//SelectExample()
	//AtomicCounterDuringUpdates()
	//AtomicCounter()
	// MutexExample()
	ProcessManager()
	//ChannelWorkerPoolExample()
	//FanOutExample()
	// DoneChannelExample()
	//CancellationExample()
	//SerialProcessCities()
}

// func LaunchCountdown(launchTime int) {
// 	for launchTime > 0 {
// 		fmt.Println(launchTime)
// 		launchTime--
// 		time.Sleep(1 * time.Second)
// 	}
// }
