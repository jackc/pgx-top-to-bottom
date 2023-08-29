package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

type CircleF32 struct {
	X float32
	Y float32
	R float32
}

func (c *CircleF32) ScanCircle(v pgtype.Circle) error {
	if !v.Valid {
		return fmt.Errorf("cannot scan null circle")
	}

	c.X = float32(v.P.X)
	c.Y = float32(v.P.Y)
	c.R = float32(v.R)

	return nil
}

func (c CircleF32) CircleValue() (pgtype.Circle, error) {
	return pgtype.Circle{
		P: pgtype.Vec2{
			X: float64(c.X),
			Y: float64(c.Y),
		},
		R:     float64(c.R),
		Valid: true,
	}, nil
}

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	var c CircleF32
	err = conn.QueryRow(ctx, "select '<(3,4),2>'::circle").Scan(&c)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c.X, c.Y, c.R) // => 3 4 2
}
