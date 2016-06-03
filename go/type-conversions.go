// type-conversations.go
// 
// Program that demonstrates how to convert
// values of one type v to another type T. T(v)
// 2016 (c) zubernetes
package main

import (
	"fmt"
	"math"
)

func print(x int) {
	fmt.Println(x)
}

func add(x int, y int) int {
	z := x + y
	return z
}

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	print(5)
	var a = add(2,2)
	print(a)
	fmt.Println(x, y, z)
}
