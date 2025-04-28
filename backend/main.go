package main

import (
	"fmt"
	"log"
	"net/http"

	"database/sqql"
	"os"

	_ "github.com/lib/pq"
)

// CONSTANTS
const PORT = ":6942"

var db *sql.DB

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Only respond to GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		return
	}

	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Handle / route
	http.HandleFunc("/", helloHandler)

	// Start the server on PORT
	fmt.Printf("Server is starting on http://localhost:%s", PORT)
	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
