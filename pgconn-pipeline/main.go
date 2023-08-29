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

	pipeline := conn.StartPipeline(ctx)
	pipeline.SendQueryParams(`select 1`, nil, nil, nil, nil)
	pipeline.SendQueryParams(`select 2`, nil, nil, nil, nil)
	pipeline.SendQueryParams(`select 3`, nil, nil, nil, nil)
	err = pipeline.Sync()
	if err != nil {
		log.Fatal(err)
	}

	// Query 1

	results, err := pipeline.GetResults()
	if err != nil {
		log.Fatal(err)
	}

	rr, ok := results.(*pgconn.ResultReader)
	if !ok {
		log.Fatalf("expected ResultReader, got: %#v", results)
	}

	for rr.NextRow() {
		row := rr.Values()
		fmt.Println(string(row[0])) // 1
	}
	commandTag, err := rr.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(commandTag) // SELECT 1

	// Query 2

	results, err = pipeline.GetResults()
	if err != nil {
		log.Fatal(err)
	}

	rr, ok = results.(*pgconn.ResultReader)
	if !ok {
		log.Fatalf("expected ResultReader, got: %#v", results)
	}

	for rr.NextRow() {
		row := rr.Values()
		fmt.Println(string(row[0])) // 2
	}
	commandTag, err = rr.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(commandTag) // SELECT 1

	// Query 3

	results, err = pipeline.GetResults()
	if err != nil {
		log.Fatal(err)
	}

	rr, ok = results.(*pgconn.ResultReader)
	if !ok {
		log.Fatalf("expected ResultReader, got: %#v", results)
	}

	for rr.NextRow() {
		row := rr.Values()
		fmt.Println(string(row[0])) // 3
	}
	commandTag, err = rr.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(commandTag) // SELECT 1

	// Sync response

	results, err = pipeline.GetResults()
	if err != nil {
		log.Fatal(err)
	}

	_, ok = results.(*pgconn.PipelineSync)
	if !ok {
		log.Fatalf("expected ResultReader, got: %#v", results)
	}

	// Send more queries

	pipeline.SendQueryParams(`select 4`, nil, nil, nil, nil)
	pipeline.SendQueryParams(`select 5`, nil, nil, nil, nil)
	err = pipeline.Sync()
	if err != nil {
		log.Fatal(err)
	}

	// Query 4

	results, err = pipeline.GetResults()
	if err != nil {
		log.Fatal(err)
	}

	rr, ok = results.(*pgconn.ResultReader)
	if !ok {
		log.Fatalf("expected ResultReader, got: %#v", results)
	}

	for rr.NextRow() {
		row := rr.Values()
		fmt.Println(string(row[0])) // 4
	}
	commandTag, err = rr.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(commandTag) // SELECT 1

	// Query 5

	results, err = pipeline.GetResults()
	if err != nil {
		log.Fatal(err)
	}

	rr, ok = results.(*pgconn.ResultReader)
	if !ok {
		log.Fatalf("expected ResultReader, got: %#v", results)
	}

	for rr.NextRow() {
		row := rr.Values()
		fmt.Println(string(row[0])) // 5
	}
	commandTag, err = rr.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(commandTag) // SELECT 1

	// Sync response

	results, err = pipeline.GetResults()
	if err != nil {
		log.Fatal(err)
	}

	_, ok = results.(*pgconn.PipelineSync)
	if !ok {
		log.Fatalf("expected ResultReader, got: %#v", results)
	}

	err = pipeline.Close()
	if err != nil {
		log.Fatal(err)
	}
}
