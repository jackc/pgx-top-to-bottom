package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	batch := &pgx.Batch{}
	batch.Queue("select 'foo'").QueryRow(func(row pgx.Row) error {
		var s string
		err := row.Scan(&s)
		if err != nil {
			return err
		}
		fmt.Println(s) // => foo
		return nil
	})
	batch.Queue("select 'bar'").QueryRow(func(row pgx.Row) error {
		var s string
		err := row.Scan(&s)
		if err != nil {
			return err
		}
		fmt.Println(s) // => foo
		return nil
	})
	err = conn.SendBatch(ctx, batch).Close()
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	//
	// foo
	// bar
}
