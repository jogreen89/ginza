// for-continued.go
//
// This program demonstrates that the init and
// post statements are optional.
// 2016 (c) zubernetes
package main

import "fmt"

func main() {
    sum := 1
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
}

