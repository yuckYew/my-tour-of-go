package main

import (
    "fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func Sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, ErrNegativeSqrt(x)
    }
    var z float64 = 1.0
    var z_prev float64
    for i := 0; i < 10; i++ {
        z -= (z*z - x) / (2*z)
        fmt.Println(z)
        // Not thought well enough.
        // Buggy when calculating Sqrt(2)
        if z_prev == z{
            return z, nil
        }
        z_prev = z
    }
    return z, nil
}

func main() {
    fmt.Println(Sqrt(2))
    fmt.Println(Sqrt(-2))
}
