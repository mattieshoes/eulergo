package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sumSquare := uint64(0)
	squareSum := uint64(0)

	for i := uint64(1); i <= 100; i++ {
		squareSum += i
		sumSquare += i * i
	}
	squareSum *= squareSum
	end := time.Now()
	fmt.Println("Answer:", squareSum-sumSquare)
	fmt.Printf("Time: %v\n", end.Sub(start))
}
