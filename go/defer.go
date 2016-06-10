// defer.go
//
// Defer statement that executes when the local function returns. In
// this program the function 'Println' returns after the subsequent call
// to 'Println("hello")'.
// 2016 (c) zubernetes
package main

import "fmt"

func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
