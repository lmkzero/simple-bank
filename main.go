// Package main ...
package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmkzero/simple-bank/internal/config"
	"github.com/lmkzero/simple-bank/internal/data"
	"github.com/lmkzero/simple-bank/internal/server"
)

func main() {
	cfg, err := config.Load("./config")
	if err != nil {
		log.Fatal("load config: ", err)
	}
	pool, err := pgxpool.New(context.Background(), cfg.DBSource)
	if err != nil {
		log.Fatal("connect to db: ", err)
	}
	store := data.NewStore(pool)
	server := server.NewServer(store)
	if err := server.Init(); err != nil {
		log.Fatal("init server: ", err)
	}
	if err := server.RegisterValidator(); err != nil {
		log.Fatal("register validator: ", err)
	}
	server.RegisterService()
	if err := server.Start(cfg.ServerAddress); err != nil {
		log.Fatal("start server: ", err)
	}
}
