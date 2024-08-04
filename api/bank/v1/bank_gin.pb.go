// Code generated by github.com/varluffy/protoc-gen-go-gin. DO NOT EDIT.

package v1

import (
	context "context"
	gin "github.com/gin-gonic/gin"
	ginx "github.com/varluffy/rich/transport/http/gin/ginx"
	metadata "google.golang.org/grpc/metadata"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the varluffy/protoc-gen-go-gin package it is being compiled against.
// context.metadata.
// gin.ginx.

type BankHTTPServer interface {
	CreateAccount(context.Context, *CreateAccountReq) (*CreateAccountRsp, error)
}

func RegisterBankHTTPServer(r gin.IRouter, srv BankHTTPServer) {
	s := Bank{
		server: srv,
		router: r,
	}
	s.RegisterService()
}

type Bank struct {
	server BankHTTPServer
	router gin.IRouter
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

func (s *Bank) RegisterService() {

	s.router.Handle("POST", "/accounts", s.CreateAccount_0)

}
