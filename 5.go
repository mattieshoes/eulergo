/*
    2520 is the smallest number that can be divided by each of the numbers from
    1 to 10 without any remainder.

    What is the smallest positive number that is evenly divisible by all of the
    numbers from 1 to 20?
*/

package main

import (
	"fmt"
	"time"
)

// we can ignore the numbers from 1-10 since anything divisible by 11-20
// will automatically be divisible by 1-10

func main() {
	start := time.Now()
	var val uint64
	for val = 20; ; val += 20 {
		found := true
		for div := uint64(20); div > 10; div-- {
			if val%div != 0 {
				found = false
				break
			}
		}
		if found {
			break
		}
	}

	end := time.Now()
	fmt.Println("Answer:", val)
	fmt.Printf("Time: %v\n", end.Sub(start))
}
