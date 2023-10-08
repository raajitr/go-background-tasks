package main

import (
	"fmt"
	"time"
)

type BackroundTask struct  {
	isRunning chan bool
	doneRow chan int
}

func NewBackgroundTask(isRunning chan bool) *BackroundTask{
	return &BackroundTask{
		isRunning: isRunning,
		doneRow: make(chan int),
	}
}

func (b BackroundTask) toggleStatus(status bool) {
	b.isRunning <- status
}

func (b BackroundTask) Start() {
	b.isRunning <- true

	// Simulate a long-running background job
	fmt.Println("Background job started...")
	time.Sleep(5 * time.Second)

	defer func(){
		fmt.Println("Background job completed.")
		b.isRunning <- false
	}()

}