// struct-pointers.go
//
// A Go example of pointers to structs,
// from tour.golang.org/moretypes/4.
// 2016 (c) zubernetes
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}
