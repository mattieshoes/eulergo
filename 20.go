/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import (
	"fmt"
	"time"
    "math/big"
)

func factorial(n int64) *big.Int {
    bigN := big.NewInt(n)
    if n == 2 {
        return bigN
    }
    bigN.Mul(bigN, factorial(n-1))
    return bigN
}


func main() {
	start := time.Now()

    //calculate
    val := factorial(int64(100))

    // convert to string, then add digits
    sum := uint64(0)
    s := fmt.Sprintf("%d", val)
    for _, i := range(s) {
        sum += uint64(i - '0')
    }

    end := time.Now()
    fmt.Println("Sum:", sum)
	fmt.Printf("Time: %v\n", end.Sub(start))
}
