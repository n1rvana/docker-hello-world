package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting up...")
	http.HandleFunc("/", handler)
	fmt.Println("Start webserver")
	http.ListenAndServe(":8080", nil)
	fmt.Println("Finishing")
}

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "<html><head></head><body><h3>Hello World from a docker container...in a flynn cluster...that's been updated</h3></body>")
}
