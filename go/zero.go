// zero.go
//
// zero value:
// 
// numeric types - 0
// boolean type  - false
// strings       - ""    (empty string)
// 2016 (c) zubernetes
package main

import "fmt"

func main() {
	i := 4
	f := 5.128
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
