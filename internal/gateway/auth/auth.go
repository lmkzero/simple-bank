// Package auth 自定义鉴权中间件
package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lmkzero/simple-bank/internal/biz/token"
)

const (
	keyAuthorizationHeader   = "authorization"
	defaultAuthorizationType = "bearer"
)

// MiddlewareFunc 返回鉴权中间件实现函数
func MiddlewareFunc(manager token.Manager) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader(keyAuthorizationHeader)
		if len(header) == 0 {
			ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("authorization header is not found"))
			return
		}
		fields := strings.Fields(header)
		if len(fields) != 2 {
			ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("illegal authorization header pattern"))
			return
		}
		if strings.ToLower(fields[0]) != defaultAuthorizationType {
			ctx.AbortWithError(http.StatusUnauthorized, fmt.Errorf("unsupported authorization type"))
			return
		}
		if _, err := manager.Verify(fields[1]); err != nil {
			ctx.AbortWithError(http.StatusUnauthorized, err)
			return
		}
		ctx.Next()
	}
}
