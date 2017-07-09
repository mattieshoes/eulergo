package main

import (
	"fmt"
	"time"
)

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
