/*
    The following iterative sequence is defined for the set of positive 
    integers:

    n → n/2 (n is even)
    n → 3n + 1 (n is odd)

    Using the rule above and starting with 13, we generate the following 
    sequence:

    13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
    It can be seen that this sequence (starting at 13 and finishing at 1) 
    contains 10 terms. Although it has not been proved yet (Collatz Problem), 
    it is thought that all starting numbers finish at 1.

    Which starting number, under one million, produces the longest chain?

    NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package main

import (
	"fmt"
	"time"
)

// memory backed function which caches partial results in a map
func collatz(n uint64, solved map[uint64]uint64) uint64 {
    answer := solved[n]
    if(answer > 0) {
        return answer
    }
    if n % 2 == 0 {
        answer = collatz(n/2, solved) + 1
    } else {
        answer = collatz(n*3+1, solved) + 1
    }
    solved[n] = answer
    return answer
}

func main() {
	start := time.Now()
    solved := make(map[uint64]uint64)
    solved[1] = 1
    best := uint64(0)
    bestI := uint64(0)
    for i := uint64(1); i < 1e6; i++ {
        answer := collatz(i, solved)
        if answer > best {
            best = answer
            bestI = i
        }
    }
	end := time.Now()
    fmt.Printf("Best: %d (%d)\n", bestI, best)
	fmt.Printf("Time: %v\n", end.Sub(start))
}
