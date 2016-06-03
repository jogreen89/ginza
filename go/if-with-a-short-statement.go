// if-with-a-short-statement.go
//
// A program that demonstrates that a short statement
// can be executed before condition evaluation.
// 2016 (c) zubernetes
package main

import (
    "fmt"
    "math"
)

func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
        return v
    }
    return lim
}

func main() {
    fmt.Println(pow(3, 2, 10), pow(3, 3, 20),)
}
