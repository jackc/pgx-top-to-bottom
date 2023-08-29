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

	rr := conn.ExecParams(ctx,
		`select n, n::float * 1.1, n::float * 1.1 from generate_series(1, $1) n;`,
		[][]byte{[]byte("3")}, // param values
		nil,                   // param OIDs (let PostgreSQL infer them)
		nil,                   // param formats (defaults to text)
		[]int16{0, 0, 1},      // result formats (text, text, binary)
	)

	for rr.NextRow() {
		row := rr.Values()
		fmt.Printf(
			"%v (%v), %v (%v), %v\n",
			row[0], string(row[0]), row[1], string(row[1]), row[2],
		)
	}
	commandTag, err := rr.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(commandTag)

	// Output:
	//
	// [49] (1), [49 46 49] (1.1), [63 241 153 153 153 153 153 154]
	// [50] (2), [50 46 50] (2.2), [64 1 153 153 153 153 153 154]
	// [51] (3), [51 46 51 48 48 48 48 48 48 48 48 48 48 48 48 48 48 51] (3.3000000000000003), [64 10 102 102 102 102 102 103]
	// SELECT 3
}
