package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    var val uint64
    for val = 20;; val += 20 {
        found := true
        for div := uint64(20); div > 10; div-- {
            if val % div != 0 {
                found = false
                break
            }
        }
        if found {
            break
        }
    }

    end := time.Now()
    fmt.Println("Answer:", val);
    fmt.Printf("Time: %v\n", end.Sub(start))
}
