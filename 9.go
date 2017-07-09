/*
    A Pythagorean triplet is a set of three natural numbers, a < b < c, for 
    which,

    a^2 + b^2 = c^2
    For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

    There exists exactly one Pythagorean triplet for which a + b + c = 1000.
    Find the product abc.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var a, b, c uint64
	found := false

	for a = 1; a < 333; a++ {
		for b = a + 1; 1000 - a - b > b; b++ {
			c = 1000 - a - b
			if a*a+b*b == c*c {
				found = true
				break
			}

		}
		if found {
			break
		}
	}

	end := time.Now()
	fmt.Println(a, b, c)
	fmt.Println("Product:", a*b*c)
	fmt.Printf("Time: %v\n", end.Sub(start))
}
