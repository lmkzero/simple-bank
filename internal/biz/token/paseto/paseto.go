// Package paseto 提供Paseto Token实现
package paseto

import (
	"fmt"
	"time"

	"github.com/lmkzero/simple-bank/internal/biz/token"
	"github.com/o1egl/paseto/v2"
	"golang.org/x/crypto/chacha20poly1305"
)

// PasetoManager Paseto Token管理
type PasetoManager struct {
	paseto *paseto.V2
	secret []byte
}

// New 工厂方法
func New(secret string) (token.Manager, error) {
	if len(secret) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be %d characters", chacha20poly1305.KeySize)
	}
	return &PasetoManager{
		paseto: paseto.NewV2(),
		secret: []byte(secret),
	}, nil
}

func (m *PasetoManager) Create(userName string, duration time.Duration) (string, error) {
	payload, err := token.NewPayload(userName, duration)
	if err != nil {
		return "", err
	}
	return m.paseto.Encrypt(m.secret, payload, nil)
}

func (m *PasetoManager) Verify(pasetoToken string) (*token.Payload, error) {
	payload := &token.Payload{}
	if err := m.paseto.Decrypt(pasetoToken, m.secret, payload, nil); err != nil {
		return nil, token.ErrInvalidToken
	}
	if err := payload.Valid(); err != nil {
		return nil, err
	}
	return payload, nil
}
