/*
    A perfect number is a number for which the sum of its proper divisors is 
    exactly equal to the number. For example, the sum of the proper divisors 
    of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect 
    number.

    A number n is called deficient if the sum of its proper divisors is less 
    than n and it is called abundant if this sum exceeds n.

    As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest
    number that can be written as the sum of two abundant numbers is 24. By 
    mathematical analysis, it can be shown that all integers greater than 28123
    can be written as the sum of two abundant numbers. However, this upper 
    limit cannot be reduced any further by analysis even though it is known 
    that the greatest number that cannot be expressed as the sum of two 
    abundant numbers is less than this limit.

    Find the sum of all the positive integers which cannot be written as the
    sum of two abundant numbers.
*/

package eulergo

import (
	"fmt"
    "math"
)

func divisors(n uint64) []uint64 {
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

func sumDivisors(n uint64) uint64 {
    list := divisors(n)
    sum := uint64(0)
    for _, val := range(list) {
        sum += val
    }
    return sum
}

func Solution23() {
    // Create abundant number list
    abundant := make([]uint64, 0, 0)
    for i := uint64(0); i < 28123; i++ {
        if sumDivisors(i) > i {
            abundant = append(abundant, i)
        }
    }

    // Mark all the ones we can make as true
    var canMake [28124]bool
    for a := range(abundant) {
        for b := a; abundant[a]+abundant[b] < 28124; b++ {
            canMake[abundant[a]+abundant[b]] = true
        }
    }

    // Sum the falses
    sum := uint64(0)
    for i, val := range(canMake) {
        if !val {
            sum += uint64(i)
        }
    }
    fmt.Println("Sum:", sum)
}
