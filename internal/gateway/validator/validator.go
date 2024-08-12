// Package validator 自定义校验插件
package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/lmkzero/simple-bank/internal/gateway/validator/currency"
)

var vs = make(map[string]validator.Func)

// Init 初始化校验插件
func Init() {
	vs["currency"] = currency.Validate
}

// CustomValidators
func CustomValidators() map[string]validator.Func {
	return vs
}
