package main

import "fmt"

type Node struct {
	val int
	next *Node // Pointer to next Node
}

func build() {
	last := Node { val: 4 }
	ptr := &last
	next := Node { 2, &last } 
	fmt.Println(last)
	fmt.Println(&last)
	fmt.Println(ptr)
	fmt.Println(next)
	fmt.Println(&next)
}

func main() {
	fmt.Println("Hello World!")
	build()
}
