package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Change "*" to "http://localhost:3000" for more security
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Fprintf(w, "Initial setup for golang apis, containerised with hot-reloading!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
