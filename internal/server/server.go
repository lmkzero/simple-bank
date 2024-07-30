// Package server 服务运行层
package server

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/lmkzero/simple-bank/api/bank/v1"
	"github.com/lmkzero/simple-bank/internal/data"
	"github.com/lmkzero/simple-bank/internal/service"
)

// Server 服务实例
type Server struct {
	service *service.BankService
	router  *gin.Engine
}

// NewServer 工厂方法
func NewServer(store *data.Store) *Server {
	service := service.NewBankService(store)
	return &Server{
		service: service,
		router:  newRouter(service),
	}
}

func newRouter(service *service.BankService) *gin.Engine {
	router := gin.Default()
	router.POST("/accounts", createAccountHandler(service))
	return router
}

func createAccountHandler(service *service.BankService) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		req := &v1.CreateAccountReq{}
		if err := ctx.BindJSON(req); err != nil {
			ctx.JSON(http.StatusBadRequest, httpErrorMessage(err))
			return
		}
		if err := req.Validate(); err != nil {
			ctx.JSON(http.StatusBadRequest, httpErrorMessage(err))
			return
		}
		rsp, err := service.CreateAccount(context.Background(), req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, httpErrorMessage(err))
			return
		}
		ctx.JSON(http.StatusOK, rsp)
	}
}

func httpErrorMessage(err error) gin.H {
	return gin.H{"error": err.Error()}
}

// Start 启动服务
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}
