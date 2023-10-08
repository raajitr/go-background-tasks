package main

import (
	"fmt"
	"net/http"
)

var isRunning = make(chan bool, 1)

func main() {
	isRunning <- false
	backgroundTask := NewBackgroundTask(isRunning)
	fmt.Println("BG INITIATED")
    // Define handler functions for the "start" and "stats" endpoints
    http.HandleFunc("/start", startHandler(*backgroundTask))
    http.HandleFunc("/stats", statsHandler)

    // Start the HTTP server on port 8080
    fmt.Println("Server is running on :8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func startHandler(backgorundTask BackroundTask) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {

		if <-isRunning {
			fmt.Fprintln(w, "It's already running")
			return
		}
		// Handle the "start" endpoint
		fmt.Fprintln(w, "Starting the job")

		go backgorundTask.Start()
	}
}
func statsHandler(w http.ResponseWriter, r *http.Request) {
    // Handle the "stats" endpoint
    fmt.Fprintln(w, "This is the 'stats' endpoint!")
}
