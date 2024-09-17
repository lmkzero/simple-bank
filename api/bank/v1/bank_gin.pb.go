package v1

import (
	context "context"

	gin "github.com/gin-gonic/gin"
	"github.com/lmkzero/simple-bank/internal/deps"
	"github.com/lmkzero/simple-bank/internal/gateway/auth"
	ginx "github.com/varluffy/rich/transport/http/gin/ginx"
	metadata "google.golang.org/grpc/metadata"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the varluffy/protoc-gen-go-gin package it is being compiled against.
// context.metadata.
// gin.ginx.

type BankHTTPServer interface {
	CreateAccount(context.Context, *CreateAccountReq) (*CreateAccountRsp, error)

	CreateUser(context.Context, *CreateUserReq) (*CreateUserRsp, error)

	GetAccount(context.Context, *GetAccountReq) (*GetAccountRsp, error)

	ListAccounts(context.Context, *ListAccountsReq) (*ListAccountsRsp, error)

	Login(context.Context, *LoginReq) (*LoginRsp, error)

	Transfer(context.Context, *TransferReq) (*TransferRsp, error)
}

func RegisterBankHTTPServer(r gin.IRouter, srv BankHTTPServer, deps *deps.Info) {
	s := Bank{
		server: srv,
		router: r,
	}
	s.RegisterService(deps)
}

type Bank struct {
	server BankHTTPServer
	router gin.IRouter
}

func (s *Bank) CreateUser_0(ctx *gin.Context) {
	var in CreateUserReq

	if err := ginx.ShouldBind(ctx, &in); err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx.Request.Context(), md)
	out, err := s.server.(BankHTTPServer).CreateUser(newCtx, &in)
	if err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}

	ginx.Response(ctx, out)
}

func (s *Bank) Login_0(ctx *gin.Context) {
	var in LoginReq

	if err := ginx.ShouldBind(ctx, &in); err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx.Request.Context(), md)
	out, err := s.server.(BankHTTPServer).Login(newCtx, &in)
	if err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}

	ginx.Response(ctx, out)
}

func (s *Bank) CreateAccount_0(ctx *gin.Context) {
	var in CreateAccountReq

	if err := ginx.ShouldBind(ctx, &in); err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx.Request.Context(), md)
	out, err := s.server.(BankHTTPServer).CreateAccount(newCtx, &in)
	if err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}

	ginx.Response(ctx, out)
}

func (s *Bank) GetAccount_0(ctx *gin.Context) {
	var in GetAccountReq

	if err := ginx.ShouldBindUri(ctx, &in); err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}

	if err := ginx.ShouldBind(ctx, &in); err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx.Request.Context(), md)
	out, err := s.server.(BankHTTPServer).GetAccount(newCtx, &in)
	if err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}

	ginx.Response(ctx, out)
}

func (s *Bank) ListAccounts_0(ctx *gin.Context) {
	var in ListAccountsReq

	if err := ginx.ShouldBind(ctx, &in); err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx.Request.Context(), md)
	out, err := s.server.(BankHTTPServer).ListAccounts(newCtx, &in)
	if err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}

	ginx.Response(ctx, out)
}

func (s *Bank) Transfer_0(ctx *gin.Context) {
	var in TransferReq

	if err := ginx.ShouldBind(ctx, &in); err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}
	md := metadata.New(nil)
	for k, v := range ctx.Request.Header {
		md.Set(k, v...)
	}
	newCtx := metadata.NewIncomingContext(ctx.Request.Context(), md)
	out, err := s.server.(BankHTTPServer).Transfer(newCtx, &in)
	if err != nil {
		ginx.ErrorResponse(ctx, err)
		return
	}

	ginx.Response(ctx, out)
}

func (s *Bank) RegisterService(deps *deps.Info) {
	s.router.Handle("POST", "/users", s.CreateUser_0)
	s.router.Handle("POST", "/users/login", s.Login_0)
	authGroup := s.router.Group("/").Use(auth.MiddlewareFunc(deps.Token))
	authGroup.Handle("POST", "/accounts", s.CreateAccount_0)
	authGroup.Handle("GET", "/accounts/:id", s.GetAccount_0)
	authGroup.Handle("GET", "/accounts", s.ListAccounts_0)
	authGroup.Handle("POST", "/transfer", s.Transfer_0)
}
