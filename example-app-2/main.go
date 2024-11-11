package main

import (
	"fmt"
	"net/http"
)

const version = "1.0.0"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from example-app-2! version %s", version)
	})
	http.ListenAndServe(":8080", nil)
}
