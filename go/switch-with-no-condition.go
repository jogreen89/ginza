// switch-with-no-condition.go
//
// As it says, a switch construct with no condition, 
// which is a good way to write long if-then-else chains.
// 2016 (c) zubernetes
package main

import (
    "fmt"
    "time"
)

func main() {
    t := time.Now()

    switch {
    case t.Hour() < 12:
        fmt.Println("Good morning!")
    case t.Hour() < 17:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Good evening.")
    }
}
