// exercise-http-handlers.go
//
// Example of using http handlers for localhost
// 2016 (c) zubernetes
package main

import (
	"log"
	"net/http"
)

func main() {
	// your http.Handle calls here
	log.Fatal(http.ListenAndServe("localhost:4000",nil))
}
