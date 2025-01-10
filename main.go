package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fprintf, err := fmt.Fprintf(w, "Hello, Dockerized Golang!")
	if err != nil {
		return
	}
	_ = fprintf
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on port 8080...")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		return
	}
}
