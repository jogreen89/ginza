// imports.go
//
// Grouping imports into a factored import statement.
// 2016 (c) zubernetes
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Now you have %g problems", math.Nextafter(2, 3))
}
