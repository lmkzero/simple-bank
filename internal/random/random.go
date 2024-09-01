// Package random 自定义random工具
package random

import (
	"strings"

	"golang.org/x/exp/rand"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// Int64 随机范围整数
func Int64(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// NString 长度n的随机字符串
func NString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}
