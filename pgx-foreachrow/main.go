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

	rows, _ := conn.Query(ctx, "select generate_series(1, 5)")
	var n int32
	_, err = pgx.ForEachRow(rows, []any{&n}, func() error {
		fmt.Println(n)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	// Output:
	//
	// 1
	// 2
	// 3
	// 4
	// 5
}
