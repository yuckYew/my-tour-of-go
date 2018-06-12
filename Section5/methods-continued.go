package main

import (
    "fmt"
    //"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

func main()  {
    // f := MyFloat(-math.Sqrt2)
    f := MyFloat(-2.0)
    fmt.Println(f.Abs())
}
