package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Background struct {
	isRunning chan bool
	doneRow   chan int
	db        *pgxpool.Pool
}

func NewBackgroundTask(db *pgxpool.Pool) *Background {
	ret := &Background{
		isRunning: make(chan bool, 1),
		doneRow:   make(chan int, 1),
		db:        db,
	}

	ret.isRunning <- false

	return ret
}

func (b Background) toggleStatus(status bool) {
	b.isRunning <- status
}

func (b Background) updateValues() error {
	query := "SELECT id FROM test.Person FOR UPDATE SKIP LOCKED"
	rows, err := b.db.Query(context.Background(), query)
	fmt.Println(rows.Next(), "Line 1")

	if err != nil {
		fmt.Println(err, "Line 1")
		return err
	}

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			fmt.Println(err, "Line 2")
			return err
		}

		// Update the value for the current row
		updateQuery := "UPDATE test.Person SET LastUpdated = $1 WHERE id = $2"
		_, err := b.db.Exec(context.Background(), updateQuery, time.Now(), id)
		if err != nil {
			fmt.Println(err, "Line 3")
			return err
		}
		fmt.Println(id)
		
		// don't really care if the message is being listened or not.
		select {
			case b.doneRow <- id:
				continue; 
			default:
				continue;
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err, "Line 4")
		return err
	}

	return nil
}

func (b Background) Start() {
	b.isRunning <- true

	// Simulate a long-running background job
	fmt.Println("Background job started...")
	b.updateValues()

	curr_status := <-b.isRunning
	b.isRunning <- !curr_status
	fmt.Println("Background job completed.")

}
