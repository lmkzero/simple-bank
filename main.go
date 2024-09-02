// Package main ...
package main

import (
	"log"

	"github.com/lmkzero/simple-bank/internal/config"
	"github.com/lmkzero/simple-bank/internal/deps"
	"github.com/lmkzero/simple-bank/internal/server"
)

func main() {
	cfg, err := config.Load("./config")
	if err != nil {
		log.Fatal("load config: ", err)
	}
	deps, err := deps.New(cfg)
	if err != nil {
		log.Fatal("init deps: ", err)
	}
	server := server.NewServer(deps)
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
