package main

import (
    "fmt"
    "time"
)

func main() {
    start := time.Now()
    var sum int
    for i := 1; i < 1000; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum += i
        }
    }
    end := time.Now()
    fmt.Println("Sum:", sum)
    fmt.Printf("Time: %v\n", end.Sub(start))
}
