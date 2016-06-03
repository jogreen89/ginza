// exported-names.go
//
// A review of exports. Capital letters are exported names,
// lower-case letters are not exported. (Unexported names 
// cannot be accessed when a package is imported.
// 2016 (c) zubernetes
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Pi)
}
