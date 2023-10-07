package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
    // Define handler functions for the "start" and "stats" endpoints
    http.HandleFunc("/start", startHandler)
    http.HandleFunc("/stats", statsHandler)

    // Start the HTTP server on port 8080
    fmt.Println("Server is running on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func startHandler(w http.ResponseWriter, r *http.Request) {
    // Handle the "start" endpoint
    fmt.Fprintln(w, "Welcome to the 'start' endpoint!")

	go backgroundJob()
}

func statsHandler(w http.ResponseWriter, r *http.Request) {
    // Handle the "stats" endpoint
    fmt.Fprintln(w, "This is the 'stats' endpoint!")
}


func backgroundJob() {
    // Simulate a long-running background job
    fmt.Println("Background job started...")
    time.Sleep(5 * time.Second)
    fmt.Println("Background job completed.")
}