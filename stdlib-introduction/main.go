package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var n int32
	err = db.QueryRow("select 42").Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)
}
