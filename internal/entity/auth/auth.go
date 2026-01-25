// Package auth 用户鉴权逻辑
package auth

import "golang.org/x/crypto/bcrypt"

// HashPassword 返回用户密码hash串
func HashPassword(password string) (string, error) {
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hp), nil
}

// IsPasswordValid 用户密码是否有效
func IsPasswordValid(password, hashedPassword string) (bool, error) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}
