package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "github.com/mattieshoes/eulergo/solutions"
    "time"
)

func main() {
    solution := map[int]func() {
        1 : eulergo.Solution1,
        2 : eulergo.Solution2,
        3 : eulergo.Solution3,
        4 : eulergo.Solution4,
        5 : eulergo.Solution5,
        6 : eulergo.Solution6,
        7 : eulergo.Solution7,
        8 : eulergo.Solution8,
        9 : eulergo.Solution9,
        10 : eulergo.Solution10,
        11 : eulergo.Solution11,
        12 : eulergo.Solution12,
        13 : eulergo.Solution13,
        14 : eulergo.Solution14,
        15 : eulergo.Solution15,
        16 : eulergo.Solution16,
        //17 : eulergo.Solution17,
        18 : eulergo.Solution18,
        19 : eulergo.Solution19,
        20 : eulergo.Solution20,
        21 : eulergo.Solution21,
        //22 : eulergo.Solution22,
        23 : eulergo.Solution23,
        24 : eulergo.Solution24,
        25 : eulergo.Solution25,
    }
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Print("eulergo> ")
        scanner.Scan()
        s := scanner.Text()
        i, _ := strconv.Atoi(s)

        if i > 0 {
            if val, ok := solution[i]; ok {
                start := time.Now()
                val()
                end := time.Now()
                fmt.Println("Time:", end.Sub(start))
            } else {
                fmt.Println("Solution", i, "doesn't exist")
            }
        } else {
            switch s {
                case "q":
                    return
            }
        }


    }
}
