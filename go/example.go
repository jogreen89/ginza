// example.go
//
// A Go implementation of a simple web server.
// (c) 2016 zubernetes

package main

import (
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/"))))
}
