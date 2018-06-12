package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an intA
func fibonacci() func() int{
    a := 0
    b := 1
    fibo := 0
    return func() int{
        fibo = a
        a = a + b
        b = fibo
        return fibo
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
