package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// rand.Intn(..) Random number.
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("This is another line", rand.Intn(15))
}
