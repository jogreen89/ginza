// variables-with-initializers.go
//
// var initialization does not have to include a 
// type if 'var' is included.
// 2016 (c) zubernetes
package main

import "fmt"

var i, j int = 1, 2

func main() {
    var c, python, java = true, false, "no!"
    fmt.Println(i, j, c, python, java)
}
