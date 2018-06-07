package main

import (
    "fmt"
)

func Sqrt(x float64) float64 {
    var z float64 = 1.0
    var z_prev float64
    for i := 0; i < 10; i++ {
        z -= (z*z - x) / (2*z)
        fmt.Println(z)
        // Not thought well enough.
        // Buggy when calculating Sqrt(2)
        if z_prev == z{
            return z
        }
        z_prev = z
    }
    return z
}

func main() {
    fmt.Println(Sqrt(2))
}
