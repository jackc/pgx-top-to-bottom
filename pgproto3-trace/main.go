package main

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgproto3"
)

func main() {
	ctx := context.Background()

	conn, err := pgconn.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	conn.Frontend().Trace(os.Stdout, pgproto3.TracerOptions{
		SuppressTimestamps: true,
		RegressMode:        true,
	})

	mrr := conn.Exec(ctx, `
		select 'Hello, world';
		select 'Goodbye, world';`,
	)
	err = mrr.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 	F	Query	58	 "
	// 	select 'Hello, world';
	// 	select 'Goodbye, world';"
	// B	RowDescription	34	 1 "?column?" 0 0 25 -1 -1 0
	// B	DataRow	23	 1 12 'Hello, world'
	// B	CommandComplete	14	 "SELECT 1"
	// B	RowDescription	34	 1 "?column?" 0 0 25 -1 -1 0
	// B	DataRow	25	 1 14 'Goodbye, world'
	// B	CommandComplete	14	 "SELECT 1"
	// B	ReadyForQuery	6	 I

	rr := conn.ExecParams(ctx,
		`select n, n::float * 1.1, n::float * 1.1 from generate_series(1, $1) n;`,
		[][]byte{[]byte("3")}, // param values
		nil,                   // param OIDs (let PostgreSQL infer them)
		nil,                   // param formats (defaults to text)
		[]int16{0, 0, 1},      // result formats (text, text, binary)
	)
	_, err = rr.Close()
	if err != nil {
		log.Fatal(err)
	}

	// F	Parse	80	 "" "select n, n::float * 1.1, n::float * 1.1 from generate_series(1, $1) n;" 0
	// F	Bind	24	 "" "" 0 1 '3' 3 0 0 1
	// F	Describe	7	 P ""
	// F	Execute	10	 "" 0
	// F	Sync	5
	// B	ParseComplete	5
	// B	BindComplete	5
	// B	RowDescription	81	 3 "n" 0 0 23 4 -1 0 "?column?" 0 0 701 8 -1 0 "?column?" 0 0 701 8 -1 1
	// B	DataRow	31	 3 1 '1' 3 '1.1' 8 '?\xf1\x99\x99\x99\x99\x99\x9a'
	// B	DataRow	31	 3 1 '2' 3 '2.2' 8 '@\x1\x99\x99\x99\x99\x99\x9a'
	// B	DataRow	46	 3 1 '3' 18 '3.3000000000000003' 8 '@\xafffffg'
	// B	CommandComplete	14	 "SELECT 3"
	// B	ReadyForQuery	6	 I
}
