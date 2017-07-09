/*
    A permutation is an ordered arrangement of objects. For example, 3124 is 
    one possible permutation of the digits 1, 2, 3 and 4. If all of the 
    permutations are listed numerically or alphabetically, we call it 
    lexicographic order. The lexicographic permutations of 0, 1 and 2 are:

    012   021   102   120   201   210

    What is the millionth lexicographic permutation of the digits 
    0, 1, 2, 3, 4, 5, 6, 7, 8 and 9?
*/

// We don't need to enumerate all of them because there are n! permutations of 
// n items.  So there are 9! permutations with 0 as the first number, another 
// 9! with 1 as the first number, etc.
// so we recursively solve for each number

package eulergo

import "fmt"

func fact(n int) int {
    prod := 1
    for ;n > 1; n-- {
        prod *= n
    }
    return prod
}

func lexOrder(nth int, choices, current []int) {
    // print when we have used all the choices
    if len(choices) == 0 {
        fmt.Println(current)
        return
    }

    // calculate which choices we're going to skip
    consumes := fact(len(choices)-1)
    count := 0
    for nth >= consumes {
        count++
        nth -= consumes
    }

    // add the winning choice to our current slice
    current = append(current, choices[count])

    // construct a new list of choices minus the choice we used
    newChoices := make([]int, len(choices) - 1, len(choices) - 1)
    for i := range(newChoices) {
        if i >= count {
            newChoices[i] = choices[i+1]
        } else {
            newChoices[i] = choices[i]
        }
    }
    // call lexOrder on the next digit
    lexOrder(nth, newChoices, current)
}

func Solution24() {
    choices := make([]int, 10, 10)
    current := make([]int, 0, 0)
    for i := range(choices) {
        choices[i] = i
    }

    // the first ordering is the 0th, so we want the 1,000,000th ordering 
    // has index 999,999
    lexOrder(1e6 - 1, choices, current)
}
