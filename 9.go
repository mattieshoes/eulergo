package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	var a, b, c uint64
	found := false

	for a = 1; a < 334; a++ {
		for b = a + 1; b < 500; b++ {
			c = 1000 - a - b
			if c <= b {
				break
			}

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
