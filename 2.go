/*
    Each new term in the Fibonacci sequence is generated by adding the previous 
    two terms. By starting with 1 and 2, the first 10 terms will be:

    1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

    By considering the terms in the Fibonacci sequence whose values do not 
    exceed four million, find the sum of the even-valued terms.
*/

package main

import (
	"fmt"
	"time"
)

// every third number is even, but this solution doesn't check such things

func main() {
	start := time.Now()
	var a, b, sum int64 = 1, 2, 0
	for b < 4e6 {
		if b%2 == 0 {
			sum += b
		}
		b = a + b
		a = b - a
	}
	end := time.Now()
	fmt.Println("Sum:", sum)
	fmt.Printf("Time: %v\n", end.Sub(start))
}
