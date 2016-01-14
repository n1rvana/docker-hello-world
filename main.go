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
	fmt.Fprintf(rw, "<h3>Hello World from a docker container...</h3>")
}
