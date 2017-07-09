/*
*/

package main

import (
	"fmt"
	"time"
)


func main() {
	start := time.Now()
	end := time.Now()
	fmt.Printf("Time: %v\n", end.Sub(start))
}
