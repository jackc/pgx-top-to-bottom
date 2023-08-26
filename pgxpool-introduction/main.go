package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	var n int32
	err = dbpool.QueryRow(ctx, "select $1::int", 42).Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n) // => 42
}
