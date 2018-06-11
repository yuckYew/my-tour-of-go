package main

import (
    "fmt"
    "math"
)

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {
        f := float64(i)
        fmt.Printf("2**%d = %f check = %d\n", i, math.Pow(2,f), v)
    }
}
