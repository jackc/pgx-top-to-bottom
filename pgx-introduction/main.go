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

	var n int32
	err = conn.QueryRow(ctx, "select $1::int", 42).Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n) // => 42
}
