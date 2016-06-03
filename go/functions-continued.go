// functions-continued.go
//
// Omitting the type from two or more consecutive named 
// function parameters. More function practice.
// 2016 (c) zubernetes
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
