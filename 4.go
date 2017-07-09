/*
     A palindromic number reads the same both ways. The largest palindrome made
     from the product of two 2-digit numbers is 9009 = 91 × 99.

     Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"time"
)

// reverses a string because those assholes dont have .reverse() :-(
func reverse(s string) string {
	rune := make([]rune, len(s))
	n := 0
	for _, r := range s {
		rune[n] = r
		n++
	}
	rune = rune[0:n]

	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	return string(rune)
}

// returns true if number is palindromic in base 10
func is_palindrome(n uint64) bool {
	s := fmt.Sprintf("%d", n)
	if s == reverse(s) {
		return true
	}
	return false
}

func main() {
	start := time.Now()
	max := uint64(0)
	for a := uint64(999); a >= 100; a-- {
		for b := uint64(999); b >= 100; b-- {
			product := a * b
			if product <= max {
				break
			}
			if is_palindrome(product) {
				max = product
			}
		}
	}
	end := time.Now()
	fmt.Println("Largest:", max)
	fmt.Printf("Time: %v\n", end.Sub(start))
}
