package main

import (
	"fmt"
	"net/http"
)

func main() {
	db_conn := DBClient()
	backgroundTask := NewBackgroundTask(db_conn)
	fmt.Println("BG INITIATED")

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
			w.WriteHeader(429)
			fmt.Fprintln(w, "It's already running")
			b.isRunning <- true // keep it true since we enqued it already
			return
		}

		go b.Start()

		w.WriteHeader(202)
		fmt.Fprintln(w, "Starting the job")
	}
}
func statsHandler(b Background) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var currentRows int

		select {
			case currentRows = <-b.doneRow:
				fmt.Println(currentRows)
			default:
				fmt.Println("Probably zero")
				currentRows = 0
		}
		

		if currentRows == 0 {
			w.WriteHeader(412)
			fmt.Fprintln(w, "No Background Process Started")
			return
		}
		msg := fmt.Sprintf("Current Completed Rows: %d", currentRows)
		
		w.WriteHeader(200)
		fmt.Fprintln(w, msg)
	}
}
