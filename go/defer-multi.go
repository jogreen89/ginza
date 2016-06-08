// defer-multi.go
//
// Stacking defers are executed in LIFO order.
// 2016 (c) zubernetes
package main

import "fmt"

func main() {
    fmt.Println("counting")

    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }

    fmt.Println("done")
}

