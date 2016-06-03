// for-is-gos-while.go
//
// Dropped the semicolons. For is Go's "while"
// 2016 (c) zubernetes
package main

import "fmt"

func main() {
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
}

