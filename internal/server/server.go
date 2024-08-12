// Package server 服务运行层
package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	v1 "github.com/lmkzero/simple-bank/api/bank/v1"
	"github.com/lmkzero/simple-bank/internal/data"
	gv "github.com/lmkzero/simple-bank/internal/gateway/validator"
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

// Init 初始化部分依赖
func (s *Server) Init() error {
	gv.Init()
	return nil
}

// Start 启动服务
func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

// RegisterService 注册路由handler
func (s *Server) RegisterService() {
	v1.RegisterBankHTTPServer(s.router, s.service)
}

// RegisterValidator 注册自定义校验插件
func (s *Server) RegisterValidator() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("fail to init validator")
	}
	for tag, fn := range gv.CustomValidators() {
		if err := v.RegisterValidation(tag, fn); err != nil {
			return err
		}
	}
	return nil
}
