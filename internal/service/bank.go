// Package service 服务逻辑层-实现API接口
package service

import (
	"context"

	v1 "github.com/lmkzero/simple-bank/api/bank/v1"
	"github.com/lmkzero/simple-bank/internal/data"
	"github.com/lmkzero/simple-bank/internal/data/db"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// BankService 银行账户服务
type BankService struct {
	store *data.Store
}

// NewBankService 工厂方法
func NewBankService(store *data.Store) *BankService {
	return &BankService{
		store: store,
	}
}

// CreateAccount 创建账户
func (b *BankService) CreateAccount(ctx context.Context, req *v1.CreateAccountReq) (*v1.CreateAccountRsp, error) {
	account, err := b.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    req.GetOwner(),
		Balance:  0,
		Currency: req.GetCurrency(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateAccountRsp{
		Id:       account.ID,
		Owner:    account.Owner,
		Balance:  account.Balance,
		Currency: account.Currency,
		CreateAt: timestamppb.New(account.CreatedAt.Time),
	}, nil
}
