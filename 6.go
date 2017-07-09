/*
    The sum of the squares of the first ten natural numbers is,

    1^2 + 2^2 + ... + 10^2 = 385
    The square of the sum of the first ten natural numbers is,

    (1 + 2 + ... + 10)^2 = 552 = 3025
    Hence the difference between the sum of the squares of the first ten 
    natural numbers and the square of the sum is 3025 − 385 = 2640.

    Find the difference between the sum of the squares of the first one hundred
    natural numbers and the square of the sum.
*/

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
