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

	var numbers []int32
	rows, err := conn.Query(ctx, "select generate_series(1, 10)")
	if err != nil {
		// Error check could be omitted as any error is included in rows.
		log.Fatal(err)
	}

	for rows.Next() {
		var n int32
		err = rows.Scan(&n)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, n)
	}

	if rows.Err() != nil {
		log.Fatal(err)
	}

	fmt.Println(numbers) // => [1 2 3 4 5 6 7 8 9 10]
}
