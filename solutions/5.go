/*
   2520 is the smallest number that can be divided by each of the numbers from
   1 to 10 without any remainder.

   What is the smallest positive number that is evenly divisible by all of the
   numbers from 1 to 20?
*/

// we can ignore the numbers from 1-10 since anything divisible by 11-20
// will automatically be divisible by 1-10

package eulergo

import "fmt"

func Solution5() {
	val := uint64(20)
	for ;;val += 20 {
		found := true
		for div := uint64(20); div > 10; div-- {
			if val%div != 0 {
				found = false
				break
			}
		}
		if found {
			break
		}
	}
	fmt.Println("Answer:", val)
}
