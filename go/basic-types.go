// basic-types.go
//
// Displaying Go's basic types.
// 2016 (c) zubernetes
package main

import (
    "fmt"
    "math/cmplx"
)

var (
    z = cmplx.Sqrt(-5 + 12i)
    ToBe  bool = false
    MaxInt uint64 = 1 << 64 - 1
)

func main() {
    const f = "%T(%v)\n"
    fmt.Printf(f, Tobe, Tobe)
    fmt.Printf(f, MaxInt, MaxInt)
    fmt.Printf(f, z, z)
}
