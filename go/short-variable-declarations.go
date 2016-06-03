// short-variable-declarations.go
//
// Using the ':=' short assignment statement (only can be used'
// inside functions. 
//
// Outside functions every statement begins with a keyword
// (var, func, etc)
// 2016 (c) zubernetes
package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	
	fmt.Println(i, j, k, c, python, java)
}
