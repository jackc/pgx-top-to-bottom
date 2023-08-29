package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

type Team struct {
	Name    string
	Players []*Player
}

type Player struct {
	Name   string
	Number int32
}

func main() {
	ctx := context.Background()

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	_, err = conn.Exec(ctx, `
		create table teams (
			name text primary key
		);

		create table players (
			team_name text references teams(name),
			name text,
			number int
		);

		insert into teams (name) values ('Bulls');
		insert into players (team_name, name, number) values
			('Bulls', 'Michael Jordan', 23),
			('Bulls', 'Scottie Pippen', 33),
			('Bulls', 'Dennis Rodman', 91);
		`)
	if err != nil {
		log.Fatal(err)
	}

	var team Team
	err = conn.QueryRow(ctx, `
		select name, (
			select array_agg(row(name, number) order by name)
			from players
			where team_name = teams.name
		) as players
		from teams
		where name = 'Bulls'`,
	).Scan(&team.Name, &team.Players)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(team.Name)
	for _, player := range team.Players {
		fmt.Println(player.Name, player.Number)
	}

	// Output:
	// Bulls
	// Dennis Rodman 91
	// Michael Jordan 23
	// Scottie Pippen 33
}
