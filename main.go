package main

import (
	"fmt"
	"net/http"
)

func main() {
	db_conn := DBClient()
	backgroundTask := NewBackgroundTask(db_conn)
	fmt.Println("BG INITIATED")
	// Define handler functions for the "start" and "stats" endpoints
	http.HandleFunc("/start", startHandler(*backgroundTask))
	http.HandleFunc("/stats", statsHandler(*backgroundTask))

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func startHandler(b Background) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if <-b.isRunning {
			fmt.Fprintln(w, "It's already running")
			return
		}
		// Handle the "start" endpoint
		fmt.Fprintln(w, "Starting the job")

		go b.Start()
	}
}
func statsHandler(b Background) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Handle the "stats" endpoint
		currentRows := <-b.doneRow
		msg := fmt.Sprintf("Current Completed Rows: %d", currentRows)

		fmt.Fprintln(w, msg)
	}
}
