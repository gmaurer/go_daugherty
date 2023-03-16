package main

import (
	"fmt"
	m "math"
	"math/rand"
	"time"
)

func MathExamples() {

	p := fmt.Println

	// We need to have a seed for our random
	// Or we will get the same random every time
	rand.Seed(time.Now().UnixMicro())
	p(rand.Intn(100))

	makeRandomFloat()
	// Finding the absolute value of a number
	p(m.Abs(-300))

	// Find the square root of a float64
	p(m.Sqrt(64))

	// Rounding a float 64
	p(m.Round(10.4))

}

func makeRandomFloat() {
	p := fmt.Println
	rand.Seed(time.Now().UnixMicro())
	p(float64(rand.Intn(100)) + rand.Float64())

}
