// Package token 提供JSON Web Token认证能力
package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

// Manager token管理接口
type Manager interface {
	Create(userName string, duration time.Duration) (string, error)
	Verify(token string) (*Payload, error)
}

// Payload token载荷信息
type Payload struct {
	ID        uuid.UUID `json:"id,omitempty"`
	UserName  string    `json:"user_name,omitempty"`
	IssuedAt  time.Time `json:"issued_at,omitempty"`
	ExpiredAt time.Time `json:"expired_at,omitempty"`
}

// NewPayload 工厂方法
func NewPayload(userName string, duration time.Duration) (*Payload, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	now := time.Now()
	return &Payload{
		ID:        id,
		UserName:  userName,
		IssuedAt:  now,
		ExpiredAt: now.Add(duration),
	}, nil
}

func (p *Payload) GetExpirationTime() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: p.ExpiredAt,
	}, nil
}

func (p *Payload) GetIssuedAt() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: p.IssuedAt,
	}, nil
}

func (p *Payload) GetNotBefore() (*jwt.NumericDate, error) {
	return &jwt.NumericDate{
		Time: p.IssuedAt,
	}, nil
}

func (p *Payload) GetIssuer() (string, error) {
	return "", nil
}

func (p *Payload) GetSubject() (string, error) {
	return "", nil
}

func (p *Payload) GetAudience() (jwt.ClaimStrings, error) {
	return jwt.ClaimStrings{}, nil
}
