// Package data 封装db方法
package data

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lmkzero/simple-bank/internal/data/db"
)

// Store 提供封装的db方法和事务
type Store struct {
	pool *pgxpool.Pool
	*db.Queries
}

// NewStore 工厂方法
// * sqlc配置中使用sql_package="pgx/v5"，所以这里使用database/sql接口会冲突
func NewStore(pool *pgxpool.Pool) *Store {
	return &Store{
		pool:    pool,
		Queries: db.New(pool),
	}
}
