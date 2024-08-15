// Package service 服务逻辑层-实现API接口
package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/jackc/pgx/v5"
	v1 "github.com/lmkzero/simple-bank/api/bank/v1"
	"github.com/lmkzero/simple-bank/internal/biz/auth"
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

// CreateUser 创建用户
func (b *BankService) CreateUser(ctx context.Context, req *v1.CreateUserReq) (*v1.CreateUserRsp, error) {
	if err := req.Validate(); err != nil {
		return nil, verr.BadRequest(http.StatusBadRequest, err.Error())
	}
	hp, err := auth.HashPassword(req.GetPassword())
	if err != nil {
		return nil, verr.InternalServer(http.StatusInternalServerError, err.Error())
	}
	user, err := b.store.CreateUser(ctx, db.CreateUserParams{
		Username:       req.GetUserName(),
		HashedPassword: hp,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserRsp{
		UserName:          user.Username,
		FullName:          user.FullName,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt.Time),
		CreateAt:          timestamppb.New(user.CreatedAt.Time),
	}, nil
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

// ListAccounts 账户列表查询
func (b *BankService) ListAccounts(ctx context.Context, req *v1.ListAccountsReq) (*v1.ListAccountsRsp, error) {
	if err := req.Validate(); err != nil {
		return nil, verr.BadRequest(http.StatusBadRequest, err.Error())
	}
	accounts, err := b.store.ListAccounts(ctx, db.ListAccountsParams{
		Limit:  int32(req.GetLimit()),
		Offset: int32(req.GetOffset()),
	})
	if err != nil {
		return nil, err
	}
	pbAccounts := make([]*v1.Account, 0, len(accounts))
	for _, account := range accounts {
		pbAccounts = append(pbAccounts, &v1.Account{
			Id:       account.ID,
			Owner:    account.Owner,
			Balance:  account.Balance,
			Currency: account.Currency,
			CreateAt: timestamppb.New(account.CreatedAt.Time),
		})
	}
	return &v1.ListAccountsRsp{
		Accounts: pbAccounts,
	}, nil
}

// Transfer 转账
func (b *BankService) Transfer(ctx context.Context, req *v1.TransferReq) (*v1.TransferRsp, error) {
	if err := req.Validate(); err != nil {
		return nil, verr.BadRequest(http.StatusBadRequest, err.Error())
	}
	ok, err := b.isAccountMatched(ctx, req.GetFromAccountId(), req.GetCurrency())
	if err != nil {
		return nil, verr.NotFound(http.StatusNotFound, err.Error())
	}
	if !ok {
		return nil, verr.BadRequest(http.StatusBadRequest, "from account is not matched")
	}
	ok, err = b.isAccountMatched(ctx, req.GetToAccountId(), req.GetCurrency())
	if err != nil {
		return nil, verr.NotFound(http.StatusNotFound, err.Error())
	}
	if !ok {
		return nil, verr.BadRequest(http.StatusBadRequest, "to account is not matched")
	}
	results, err := b.store.TransferTx(ctx, data.TransferTxParams{
		FromAccountID: req.GetFromAccountId(),
		ToAccountID:   req.GetToAccountId(),
		Amount:        req.GetAmount(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.TransferRsp{
		FromAccount: &v1.Account{
			Id:       results.FromAccount.ID,
			Owner:    results.FromAccount.Owner,
			Balance:  results.FromAccount.Balance,
			Currency: results.FromAccount.Currency,
			CreateAt: timestamppb.New(results.FromAccount.CreatedAt.Time),
		},
		ToAccount: &v1.Account{
			Id:       results.Transfer.ToAccountID,
			Owner:    results.ToAccount.Owner,
			Balance:  results.ToAccount.Balance,
			Currency: results.ToAccount.Currency,
			CreateAt: timestamppb.New(results.ToAccount.CreatedAt.Time),
		},
	}, nil
}

func (b *BankService) isAccountMatched(ctx context.Context, accountID int64, currency string) (bool, error) {
	account, err := b.store.GetAccount(ctx, accountID)
	if err != nil {
		return false, err
	}
	if account.Currency != currency {
		return false, nil
	}
	return true, nil
}
