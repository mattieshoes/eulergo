/*
    Let d(n) be defined as the sum of proper divisors of n (numbers less than n
    which divide evenly into n).
    If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair 
    and each of a and b are called amicable numbers.

    For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44,
    55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4,
    71 and 142; so d(284) = 220.

    Evaluate the sum of all the amicable numbers under 10000.
*/

package eulergo

import (
	"fmt"
    "math"
)

func divisors21(n uint64) []uint64 {
    d := make([]uint64, 0, 0)
    if n <= 1 {
        return d
    }
    d = append(d, 1)
    max := uint64(math.Sqrt(float64(n)))
    for i := uint64(2); i <= max; i++ {
        if n % i == 0 {
            d = append(d, i)
            if n / i != i {
                d = append(d, n/i)
            }
        }
    }
    return d
}

func sumDivisors21(n uint64) uint64 {
    list := divisors21(n)
    sum := uint64(0)
    for _, val := range(list) {
        sum += val
    }
    return sum
}


func Solution21() {
    var list [10001]uint64
    for i := range(list) {
        list[i] = sumDivisors21(uint64(i))
    }
    sum := uint64(0)
    for i := range(list) {
        if list[i] < uint64(i) && list[list[i]] == uint64(i) {
            sum += uint64(i) + list[i]
        }
    }
    fmt.Println("Sum:", sum)
}
