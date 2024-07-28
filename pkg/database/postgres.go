package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)



func NewPostgres(dbDNS string) *pgxpool.Pool {
	ctx := context.Background()

	conn, err := pgxpool.Connect(ctx, dbDNS)

	if err != nil {
		log.Println("failed connection to postgres db")
		panic(err)
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	if err = conn.Ping(ctx); err != nil {
		log.Println("failed ping to postgres db")
		panic(err)
	}

	return conn
}
