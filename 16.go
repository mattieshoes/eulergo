/*
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

package main

import (
	"fmt"
	"time"
    "math/big"
)

// Just uses arbitrary precision integers to calculate the actual sum, converts
// it to a string, and prints the first 10 digits
func main() {
	start := time.Now()

    // calculate 2^1000
    base := big.NewInt(int64(2))
    answer := big.NewInt(int64(1))
    for i:= 0; i < 1000; i++ {
        answer.Mul(answer, base)
    }

    // convert to string, then add digits
    sum := uint64(0)
    s := fmt.Sprintf("%d", answer)
    for _, i := range(s) {
        sum += uint64(i - '0')
    }
	
    end := time.Now()
    fmt.Println("Sum:", sum)
	fmt.Printf("Time: %v\n", end.Sub(start))
}
