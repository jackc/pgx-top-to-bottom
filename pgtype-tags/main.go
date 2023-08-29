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

	type Tag string

	var tags []Tag
	err = conn.QueryRow(ctx, "select array['foo', 'bar', 'baz']").Scan(&tags)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tags) // => [foo bar baz]
}
