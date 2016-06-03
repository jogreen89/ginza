// packages.go
//
// Discussion of packages with import paths "fmt" 
// and "math/rand". 
// 2016 (c) zubernetes
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
