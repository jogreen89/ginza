// struct-literals.go
//
// A Go example of struct literals,
// from tour.golang.org/moretypes/5.
// 2016 (c) zubernetes
package main

import "fmt"

var (
    v1 = Vertex{1, 2}  // has type Vertex
    v2 = Vertex{X: 1}  // Y:0 is implicit
    v3 = Vertex{}      // X:0 and Y:0
    p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
    fmt.Println(v1, p, v2, v3)
}
