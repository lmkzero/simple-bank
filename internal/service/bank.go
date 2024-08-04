// Package service 服务逻辑层-实现API接口
package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	v1 "github.com/lmkzero/simple-bank/api/bank/v1"
	"github.com/lmkzero/simple-bank/internal/data"
	"github.com/lmkzero/simple-bank/internal/data/db"
	verr "github.com/varluffy/rich/errcode"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// BankService 银行账户服务
type BankService struct {
	store *data.Store
}

// NewBankService 工厂方法
func NewBankService(store *data.Store) v1.BankHTTPServer {
	return &BankService{
		store: store,
	}
}

// CreateAccount 创建账户
func (b *BankService) CreateAccount(ctx context.Context, req *v1.CreateAccountReq) (*v1.CreateAccountRsp, error) {
	if err := req.Validate(); err != nil {
		return nil, verr.BadRequest(http.StatusBadRequest, err.Error())
	}
	account, err := b.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    req.GetOwner(),
		Balance:  0,
		Currency: req.GetCurrency(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateAccountRsp{
		CreatedAccount: &v1.Account{
			Id:       account.ID,
			Owner:    account.Owner,
			Balance:  account.Balance,
			Currency: account.Currency,
			CreateAt: timestamppb.New(account.CreatedAt.Time),
		},
	}, nil
}

// GetAccount 查询账户
func (b *BankService) GetAccount(ctx context.Context, req *v1.GetAccountReq) (*v1.GetAccountRsp, error) {
	if err := req.Validate(); err != nil {
		return nil, verr.BadRequest(http.StatusBadRequest, err.Error())
	}
	account, err := b.store.GetAccount(ctx, req.GetId())
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, verr.NotFound(http.StatusNotFound, err.Error())
	}
	if err != nil {
		return nil, err
	}
	return &v1.GetAccountRsp{
		Account: &v1.Account{
			Id:       account.ID,
			Owner:    account.Owner,
			Balance:  account.Balance,
			Currency: account.Currency,
			CreateAt: timestamppb.New(account.CreatedAt.Time),
		},
	}, nil
}
