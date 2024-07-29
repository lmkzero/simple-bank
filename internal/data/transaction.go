package data

import (
	"context"
	"fmt"

	"github.com/lmkzero/simple-bank/internal/data/db"
)

func (s *Store) execTransaction(ctx context.Context, f func(*db.Queries) error) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	if err := f(db.New(tx)); err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("exec transaction: %w, rollback: %w", err, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}

// TransferTx 转账事务
func (s *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResults, error) {
	results := TransferTxResults{}
	if err := s.execTransaction(ctx, func(q *db.Queries) error {
		// 新增流水记录
		transfer, err := q.CreateTransfer(
			ctx,
			db.CreateTransferParams{
				FromAccountID: arg.FromAccountID,
				ToAccountID:   arg.ToAccountID,
				Amount:        arg.Amount,
			},
		)
		if err != nil {
			return err
		}
		fromEntry, err := q.CreateEntry(
			ctx,
			db.CreateEntryParams{
				AccountID: arg.FromAccountID,
				Amount:    -arg.Amount,
			},
		)
		if err != nil {
			return err
		}
		toEntry, err := q.CreateEntry(
			ctx,
			db.CreateEntryParams{
				AccountID: arg.ToAccountID,
				Amount:    arg.Amount,
			},
		)
		if err != nil {
			return err
		}
		// 更新两个账户的余额
		newFromAccount, err := q.AddAccountBalance(ctx, db.AddAccountBalanceParams{
			Amount: -arg.Amount,
			ID:     arg.FromAccountID,
		})
		if err != nil {
			return err
		}
		newToAccount, err := q.AddAccountBalance(ctx, db.AddAccountBalanceParams{
			Amount: arg.Amount,
			ID:     arg.ToAccountID,
		})
		if err != nil {
			return err
		}
		// 赋值转账事务结果
		results.Transfer = transfer
		results.FromEntry = fromEntry
		results.ToEntry = toEntry
		results.FromAccount = newFromAccount
		results.ToAccount = newToAccount
		return nil
	}); err != nil {
		return results, err
	}
	return results, nil
}

// TransferTxParams 转账事务入参
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResults 转账事务结果
type TransferTxResults struct {
	Transfer    db.Transfer `json:"transfer"`
	FromAccount db.Account  `json:"from_account"`
	ToAccount   db.Account  `json:"to_account"`
	FromEntry   db.Entry    `json:"from_entry"`
	ToEntry     db.Entry    `json:"to_entry"`
}
