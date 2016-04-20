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
