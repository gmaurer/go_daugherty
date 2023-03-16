package main

import "fmt"

func sum(s []int, out chan int) {
	sum := 0
	for _, val := range s {
		sum += val
	}

	out <- sum
}

func ChannelBlockingExample() {
	firstNumSlice := []int{24, 25, 26, 27, -29}
	secondNumSlice := []int{34, 35, 36, -32, 31}
	out := make(chan int)

	go sum(firstNumSlice, out)
	go sum(secondNumSlice, out)

	x, y := <-out, <-out

	fmt.Printf("The two summed values written from the channel are %d and %d\n", x, y)
}

func ChannelExample() {
	numberChannel := make(chan int)

	go func() {
		numberChannel <- 12
	}()
	nums := <-numberChannel

	fmt.Println(nums)
}

func BufferedChannelExample() {
	names := make(chan string, 3)

	names <- "Paul"
	names <- "George"
	names <- "Ringo"
	fmt.Println(<-names)
	names <- "John"

	fmt.Println(<-names)
	fmt.Println(<-names)
	fmt.Println(<-names)

}

func RangeBufferedChannelExample() {

	powerpuffGirls := make(chan string, 3)

	powerpuffGirls <- "Blossom"
	powerpuffGirls <- "Bubbles"
	powerpuffGirls <- "Buttercup"

	close(powerpuffGirls)

	for p := range powerpuffGirls {
		fmt.Println(p)
	}
}

func BufferedChannelBlockingExample() {
	andysToys := make(chan string, 1)
	andysToys <- "Woody"
	fmt.Println("One of Andy's toys is", <-andysToys)
	andysToys <- "Buzz"
	fmt.Println("Another of Andy's toys is", <-andysToys)
}

func IdiomaticBufferedChannelExample() {

	const numOfPositions = 8
	players := []string{"Mike Piazza", "Fred McGriff", "Craig Biggio",
		"Matt Williams", "Barry Larkin", "Dante Bichette", "Barry Bonds", "Tony Gwynn", "Todd Hundley",
		"Jeff Bagwell", "Eric Young", "Chipper Jones", "Ozzie Smith", "Gary Sheffield"}

	startingPlayers := make(chan string, len(players))

	for _, i := range players {
		startingPlayers <- i
	}

	startingLineup := make(chan string, numOfPositions)

	for i := 0; i < numOfPositions; i++ {
		go func() {
			p := <-startingPlayers
			startingLineup <- p
		}()
	}
	close(startingPlayers)

	var roster []string
	for i := 0; i < numOfPositions; i++ {
		roster = append(roster, <-startingLineup)
	}
	close(startingLineup)

	fmt.Println(roster)

}
