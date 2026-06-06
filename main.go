package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Go!")
}

func main() {
	http.HandleFunc("/", home)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}