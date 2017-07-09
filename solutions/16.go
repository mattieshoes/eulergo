/*
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

// Just uses arbitrary precision integers to calculate the actual sum, converts
// it to a string, and prints the first 10 digits

package eulergo

import (
	"fmt"
    "math/big"
)

func Solution16() {
    
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
    fmt.Println("Sum:", sum)
}
