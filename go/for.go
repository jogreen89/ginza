// for.go
// 
// Example using the only looping construct in Go, 'for'
package main

import "fmt"

func main() {
    sum := 0

    // init : executed before the first iteration: i := 0
    // condition : evaluated before every iteration: i < 10
    // post : executed at the end of every iteration: i++
    for i := 0; i < 10; i++ {
        sum += i
    }

    fmt.Println(sum)
}
