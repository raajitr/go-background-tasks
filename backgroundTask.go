package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

)

type BackroundTask struct  {
	isRunning chan bool
	doneRow chan int
	db *pgxpool.Pool
}

func NewBackgroundTask(isRunning chan bool, db *pgxpool.Pool) *BackroundTask{
	return &BackroundTask{
		isRunning: isRunning,
		doneRow: make(chan int),
		db: db,
	}
}

func (b BackroundTask) toggleStatus(status bool) {
	b.isRunning <- status
}

func (b BackroundTask) updateValues() error {
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
		b.doneRow <- id
        if err != nil {
            fmt.Println(err, "Line 3")
			return err
        }
    }

    if err := rows.Err(); err != nil {
        fmt.Println(err, "Line 4")
		return err
    }

    return nil
}

func (b BackroundTask) Start() {
	b.isRunning <- true

	// Simulate a long-running background job
	fmt.Println("Background job started...")
	b.updateValues()

	defer func(){
		fmt.Println("Background job completed.")
		b.isRunning <- false
	}()

}