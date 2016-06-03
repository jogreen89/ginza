// web-servers.go
//
// Working with the fantastic Http package
// 2016 (c) zubernetes
package main

import (
	"fmt"
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP (
	w http.ResponseWriter,
	r *http.Request)
	fmt.Fprint(w, "Hello!")
}

func main() {
	var h hello
	err := http.ListenAndServe("localhost:4000", h)
	
	if err != nil {
		log.Fatal(err)
	}
}
