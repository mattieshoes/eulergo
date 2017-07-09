/*
   The prime factors of 13195 are 5, 7, 13 and 29.

   What is the largest prime factor of the number 600851475143 ?
*/

// just removes factors from the smallest up
// what's left over will be the largest prime factor

package eulergo

import (
	"fmt"
	"math"
)

func Solution3() {
	const num uint64 = 600851475143
	cur := num
	for {
		max := uint64(math.Sqrt(float64(cur)))
		div := uint64(2)
		found := false
		for ; div <= max; div++ {
			if cur%div == 0 {
				cur /= div
				found = true
				break
			}
		}
		if !found {
			break
		}
	}
	fmt.Printf("Largest prime factor of %d is %d\n", num, cur)
}
