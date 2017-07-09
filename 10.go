package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	start := time.Now()

	// bitfield of primes 0-1999999
	// 0 is prime, 1 is not
	var bitfield [31250]uint64

	// manually remove 0 and 1 as not prime
	bitfield[0/64] |= 1 << 0 % 64
	bitfield[1/64] |= 1 << 1 % 64

	// flip bits of multiples of all primes
	max := uint64(math.Sqrt(2000000.0))
	for i := uint64(2); i <= max; i++ {
		// retrieve value
		val := (bitfield[i/64] >> i % 64) & 1
		// if bit is 0, set all multiples to 1
		if val == 0 {
			for j := i + i; j < 2000000; j += i {
				bitfield[j/64] |= 1 << (j % 64)
			}
		}
	}

	sum := uint64(0)
	//iterate through array, converting to numbers
	for a := range bitfield {
		for b := uint64(0); b < 64; b++ {
			val := uint64(a)*64 + b
			n := (bitfield[a] >> b) & 1
			if n == 0 {
				sum += val
			}
		}
	}

	end := time.Now()
	fmt.Println("Sum:", sum)
	fmt.Printf("Time: %v\n", end.Sub(start))
}
