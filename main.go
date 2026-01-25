// Package main ...
package main

import (
	"log"

	"github.com/lmkzero/simple-bank/api/bank/v1_1"
	"github.com/lmkzero/simple-bank/internal/config"
	"github.com/lmkzero/simple-bank/internal/deps"
	"github.com/lmkzero/simple-bank/internal/gateway/auth"
	"github.com/lmkzero/simple-bank/internal/service"
	"trpc.group/trpc-go/trpc-go"
	"trpc.group/trpc-go/trpc-go/server"

	_ "trpc.group/trpc-go/trpc-filter/debuglog"
	_ "trpc.group/trpc-go/trpc-filter/recovery"
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
	trpcConfig, err := trpc.LoadConfig("./config/trpc_go.yaml")
	if err != nil {
		log.Fatalf("load trpc config: %v", err)
	}
	s := trpc.NewServerWithConfig(
		trpcConfig,
		server.WithFilter(auth.NewHandleFunc(deps.Token)),
	)
	v1_1.RegisterBankService(s.Service("trpc.mengkailiu.bank.httpv1"), service.NewBankService(deps))
	if err := s.Serve(); err != nil {
		log.Fatalf("trpc run: %v", err)
	}
}
