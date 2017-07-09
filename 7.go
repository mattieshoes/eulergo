/*
    By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
    that the 6th prime is 13.

    What is the 10 001st prime number?
*/

package main

import (
	"fmt"
	"time"
)

// just builds a list of primes until it reaches 10,001 elements

func main() {
	start := time.Now()

	primes := make([]uint64, 0)
	primes = append(primes, uint64(2))
	for v := uint64(3); ; v += 2 {
		found := false
		for _, i := range primes {
			if v%i == 0 {
				found = true
				break
			}
		}
		if !found {
			primes = append(primes, v)
			if len(primes) == 10001 {
				break
			}
		}
	}

	end := time.Now()
	fmt.Println("Answer:", primes[10000])
	fmt.Printf("Time: %v\n", end.Sub(start))
}
