package main

import (
    //"golang.org/x/tour/wc"
    "fmt"
    "strings"
)

func WordCount(s string) map[string]int {
    hist := make(map[string]int)

    words := strings.Fields(s)
    for _, w := range words{
        _, ok := hist[w]
        if ok == true {
            hist[w] += 1
        } else {
            hist[w] = 1
        }
    }
    return hist
}

func main() {
    sentence := "the quick brown fox jumped over the lazy sloth dog"
    WordCount(sentence)
    fmt.Println("hist: %p", WordCount(sentence))
    //wc.Test(WordCount)
}
