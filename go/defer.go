// defer.go
//
// Defer statement that executes when the local function returns
// 2016 (c) zubernetes
package main

import "fmt"

func main() {
    defer fmt.Println("world")

    fmt.Println("hello")
}

