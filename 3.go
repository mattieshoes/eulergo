/*
    The prime factors of 13195 are 5, 7, 13 and 29.

    What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"fmt"
	"math"
	"time"
)

// just removes factors from the smallest up
// what's left over will be the largest prime factor

func main() {
	const num uint64 = 600851475143
	start := time.Now()
	cur := num
	for {
		max := uint64(math.Sqrt(float64(cur)))
		div := uint64(2)
		found := false
		for ; div <= max; div++ {
			if cur%div == 0 {
				cur /= div
				found = true
				break
			}
		}
		if !found {
			break
		}
	}
	end := time.Now()
	fmt.Printf("Largest prime factor of %d is %d\n", num, cur)
	fmt.Printf("Time: %v\n", end.Sub(start))
}
