package main

import (
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/stdlib"
)

func main() {
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	config.OnNotice = func(c *pgconn.PgConn, n *pgconn.Notice) {
		log.Printf("PID: %d; Message: %s\n", c.PID(), n.Message)
	}

	db := stdlib.OpenDB(*config)
	defer db.Close()

	_, err = db.Exec("drop table if exists foo")
	if err != nil {
		log.Fatal(err)
	}

	// config.OnNotice prints:
	// => 2023/08/24 20:42:44 PID: 86890; Message: table "foo" does not exist, skipping
}
