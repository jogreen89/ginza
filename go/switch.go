// switch.go
//
// Demonstration of the switch statement in Golang,
// which is very C-like syntax.
// 2016 (c) zubernetes
package main

import (
    "fmt"
    "runtime"
)

func main() {
    fmt.Print("Go runs on")

    switch os := runtime.GOOS; os {
    case "darwin":
        fmt.Println("OS X")
    case "linux":
        fmt.Println("Linux.")
    default:
        // freebsd, openbsd,
        // plan9, windows...
        fmt.Printf("%s.", os)
    }
}
