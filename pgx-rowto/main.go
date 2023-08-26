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

	rows, _ := conn.Query(ctx, "select generate_series(1, 10)")
	numbers, err := pgx.CollectRows(rows, pgx.RowTo[int32])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(numbers) // => [1 2 3 4 5 6 7 8 9 10]
}
