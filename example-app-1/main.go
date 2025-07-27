package main

import (
	"fmt"
	"net/http"
)

const version = "1.0.1"

func main() {
	fmt.Println("Starting example-app-1...")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from example-app-1! version %s", version)
	})
	http.ListenAndServe(":8080", nil)
}
