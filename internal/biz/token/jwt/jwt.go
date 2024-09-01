// Package jwt 提供JSON Web Token实现
package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lmkzero/simple-bank/internal/biz/token"
)

const minSecretLength = 32

// JWTManager JWT管理
type JWTManager struct {
	secret string
}

// New 工厂方法
func New(secret string) (token.Manager, error) {
	if len(secret) < minSecretLength {
		return nil, fmt.Errorf("invalid secret length, at least %d characters", minSecretLength)
	}
	return &JWTManager{
		secret: secret,
	}, nil
}

func (m *JWTManager) Create(userName string, duration time.Duration) (string, error) {
	payload, err := token.NewPayload(userName, duration)
	if err != nil {
		return "", err
	}
	jt := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jt.SignedString([]byte(m.secret))
}

func (m *JWTManager) Verify(jwtToken string) (*token.Payload, error) {
	keyFunc := func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, token.ErrInvalidToken
		}
		return []byte(m.secret), nil
	}
	jt, err := jwt.ParseWithClaims(
		jwtToken,
		&token.Payload{},
		keyFunc,
	)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, token.ErrExpiredToken
		}
		return nil, token.ErrInvalidToken
	}
	payload, ok := jt.Claims.(*token.Payload)
	if !ok {
		return nil, token.ErrInvalidToken
	}
	return payload, nil
}
