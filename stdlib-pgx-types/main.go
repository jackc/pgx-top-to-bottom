package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgtype"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.ExecContext(ctx, `
	create temporary table t (
		id int primary key generated by default as identity,
		data int[] not null
	)`)
	if err != nil {
		log.Fatal(err)
	}

	// Query arguments are passed directly to the underlying pgx conn so there is
	// no need to implement driver.Valuer if pgx already understands the type.
	_, err = db.ExecContext(ctx,
		`insert into t (data) values ($1)`,
		[]int32{1, 2, 3},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Scanning requires the use of an adapter.
	m := pgtype.NewMap()
	var data []int32
	err = db.QueryRow("select data from t limit 1").Scan(m.SQLScanner(&data))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data) // => [1 2 3]

}
