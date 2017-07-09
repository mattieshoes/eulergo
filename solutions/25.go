/*
    The Fibonacci sequence is defined by the recurrence relation:

    Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
    Hence the first 12 terms will be:

    F1 = 1
    F2 = 1
    F3 = 2
    F4 = 3
    F5 = 5
    F6 = 8
    F7 = 13
    F8 = 21
    F9 = 34
    F10 = 55
    F11 = 89
    F12 = 144
    The 12th term, F12, is the first term to contain three digits.

    What is the index of the first term in the Fibonacci sequence to contain 
    1000 digits?
*/

package eulergo

import (
	"fmt"
    "math/big"
)

// Closure that returns the index and count of digits, starting at n=3
func Fib() func() (int, int) {
    a := big.NewInt(int64(1))
    b := big.NewInt(int64(1))
    n := 2
    return func() (int,int) {
        c := big.NewInt(0)
        c.Add(b,a)
        a = b
        b = c
        n++
        s := fmt.Sprintf("%d", c)
        return n, len(s)
    }
}

func Solution25() {
    f := Fib()
    for {
        sub, digits := f()
        if digits >= 1000 {
            fmt.Println("n = ", sub)
            break;
        }
    }
}
