// Package server 服务运行层
package server

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lmkzero/simple-bank/api/bank/v1"
	"github.com/lmkzero/simple-bank/internal/data"
	"github.com/lmkzero/simple-bank/internal/service"
)

// Server 服务实例
type Server struct {
	service v1.BankHTTPServer
	router  *gin.Engine
}

// NewServer 工厂方法
func NewServer(store *data.Store) *Server {
	return &Server{
		service: service.NewBankService(store),
		router:  gin.Default(),
	}
}

// Start 启动服务
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

// RegisterService 注册路由handler
func (s *Server) RegisterService() {
	v1.RegisterBankHTTPServer(s.router, s.service)
}
