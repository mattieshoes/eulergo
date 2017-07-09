/*
Starting in the top left corner of a 2×2 grid, and only being able to move to 
the right and down, there are exactly 6 routes to the bottom right corner.


How many such routes are there through a 20×20 grid?
*/

package main

import (
	"fmt"
	"time"
)

// problem is straight combinatorics, 40!/(20! 20!) and one could use the 
// big int library to just outright calculate it. But what's the fun in that?
// This factors everything into primes, cancels out everything it can, then 
// calculates the answer without overflowing
func main() {
	start := time.Now()

    numerator := make([]uint64, 0, 0)
    denominator := make([]uint64, 0, 0)

    // factor numerator down to primes
    for i := uint64(40); i > 20; i-- {
        numerator = append(numerator, factors(i)...)
    }
    // factor denominator down to primes
    for i := uint64(20); i > 1; i-- {
        denominator = append(denominator, factors(i)...)
    }
    // cancel out matching primes in numerator and denominator
    for i := range(numerator) {
        for j := range(denominator) {
            if numerator[i] == denominator[j] {
                numerator[i], denominator[j] = 1,1
                break;
            }
        }
    }

    answer := uint64(1)
    // now math shouldn't overflow
    for _, i := range(numerator) {
        answer *= i
    }
    for _, i := range(denominator) {
        answer /= i
    }
	end := time.Now()
    fmt.Println("Answer:", answer)
	fmt.Printf("Time: %v\n", end.Sub(start))
}

func factors(num uint64) []uint64 {
    f := make([]uint64, 0, 0)
    for i := uint64(2); i <= num/2; i++ {
        if num % i == 0 {
            f = append(f, i)
            f = append(f, factors(num/i)...)
            return f
        }
    }
    f = append(f, num)
    return f
}
