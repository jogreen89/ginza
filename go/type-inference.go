// type-inference
//
// Variables without a specified type (':=' or 'var') have
// an inferred type from the value on the right hand side.
package main

import "fmt"

func main() {
    v := 42
    j := v
    fmt.Println("v is of type %T\n", v)
    fmt.Println("j is of type %T\n", j)
}
