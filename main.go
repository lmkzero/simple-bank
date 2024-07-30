// Package main ...
package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmkzero/simple-bank/internal/data"
	"github.com/lmkzero/simple-bank/internal/server"
)

// dbSource your own db url
const dbSource = ""

func main() {
	pool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("connect to db: ", err)
	}
	store := data.NewStore(pool)
	server := server.NewServer(store)
	if err := server.Start("127.0.0.1:8080"); err != nil {
		log.Fatal("start server: ", err)
	}
}
