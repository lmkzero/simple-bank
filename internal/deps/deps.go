// Package 服务依赖
package deps

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmkzero/simple-bank/internal/biz/token"
	"github.com/lmkzero/simple-bank/internal/biz/token/paseto"
	"github.com/lmkzero/simple-bank/internal/config"
	"github.com/lmkzero/simple-bank/internal/data"
)

// Info 服务依赖项
type Info struct {
	Store *data.Store
	Token token.Manager
	Cfg   *config.AppConfig
}

// New 工厂方法
func New(cfg *config.AppConfig) (*Info, error) {
	pool, err := pgxpool.New(context.Background(), cfg.DBSource)
	if err != nil {
		return nil, fmt.Errorf("connect to db: %w", err)
	}
	store := data.NewStore(pool)
	token, err := paseto.New(cfg.TokenSecret)
	if err != nil {
		return nil, fmt.Errorf("init token manager: %w", err)
	}
	return &Info{
		Store: store,
		Token: token,
		Cfg:   cfg,
	}, nil
}
