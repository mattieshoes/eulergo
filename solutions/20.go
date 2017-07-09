/*
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

// Arbitrary precision library to calculate 100 facttorial.

package eulergo

import (
	"fmt"
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

func Solution20() {
    val := factorial(int64(100))
    sum := uint64(0)
    s := fmt.Sprintf("%d", val)
    for _, i := range(s) {
        sum += uint64(i - '0')
    }
    fmt.Println("Sum:", sum)
}
