package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgconn"
)

func main() {
	ctx := context.Background()

	conn, err := pgconn.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	mrr := conn.Exec(ctx, `
		select 'Hello, world';
		select 'Goodbye, world';`,
	)
	for mrr.NextResult() { // loop through each result set
		for mrr.ResultReader().NextRow() { // loop through each row in the result set
			row := mrr.ResultReader().Values()
			fmt.Println(string(row[0]))
		}
	}
	err = mrr.Close()
	if err != nil {
		log.Fatal(err)
	}
}
