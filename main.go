package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func main() {
	port := 8080
	fmt.Printf("Example app listening at http://localhost:%d\n", port)
	http.HandleFunc("/", hello)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
