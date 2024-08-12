// Package currency 自定义校验插件
package currency

import (
	"github.com/go-playground/validator/v10"
)

var supportedCurrency = map[string]bool{
	"USD": true,
	"EUR": true,
}

// Validate 自定义校验函数
func Validate(filedLevel validator.FieldLevel) bool {
	c, ok := filedLevel.Field().Interface().(string)
	if !ok {
		return false
	}
	if _, ok := supportedCurrency[c]; ok {
		return true
	}
	return false
}
