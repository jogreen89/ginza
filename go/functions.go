// functions.go
//
// A simple function 'add' that takes two parameters of 
// type 'int'.
// 2016 (c) zubernetes
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
