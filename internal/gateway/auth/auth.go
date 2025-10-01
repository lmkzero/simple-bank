// Package auth 自定义鉴权中间件
package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/lmkzero/simple-bank/internal/entity/token"
	"trpc.group/trpc-go/trpc-go/errs"
	"trpc.group/trpc-go/trpc-go/filter"
	thttp "trpc.group/trpc-go/trpc-go/http"
)

const (
	keyAuthorizationHeader   = "authorization"
	defaultAuthorizationType = "bearer"
)

// NewHandleFunc 工厂方法
func NewHandleFunc(manager token.Manager) filter.ServerFilter {
	return func(ctx context.Context, req any, next filter.ServerHandleFunc) (any, error) {
		httpHeader := thttp.Head(ctx).Request.Header
		if len(httpHeader) == 0 {
			return nil, errs.New(http.StatusUnauthorized, "authorization header is not found")
		}
		header := httpHeader.Get(keyAuthorizationHeader)
		fields := strings.Fields(header)
		if len(fields) != 2 {
			return nil, errs.New(http.StatusUnauthorized, "illegal authorization header pattern")
		}
		if strings.ToLower(fields[0]) != defaultAuthorizationType {
			return nil, errs.New(http.StatusUnauthorized, "unsupported authorization type")
		}
		if _, err := manager.Verify(fields[1]); err != nil {
			return nil, errs.New(http.StatusUnauthorized, err.Error())
		}
		return next(ctx, req)
	}
}
