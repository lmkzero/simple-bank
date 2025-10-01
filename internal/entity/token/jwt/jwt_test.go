// Package jwt 提供JSON Web Token实现
package jwt

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/lmkzero/simple-bank/internal/entity/token"
	"github.com/lmkzero/simple-bank/pkg/random"
	"github.com/stretchr/testify/require"
)

func TestJWTManager_Verify_No_Error(t *testing.T) {
	manager, err := New(random.NString(32))
	require.NoError(t, err)
	userName, duration := random.NString(5), time.Minute
	token, err := manager.Create(userName, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	payload, err := manager.Verify(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
	require.Equal(t, userName, payload.UserName)
	require.Equal(t, duration, payload.ExpiredAt.Sub(payload.IssuedAt))
}

func TestJWTManager_Verify_Expired(t *testing.T) {
	manager, err := New(random.NString(32))
	require.NoError(t, err)
	userName, duration := random.NString(5), -time.Minute
	token, err := manager.Create(userName, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	payload, err := manager.Verify(token)
	require.Error(t, err)
	require.Nil(t, payload)
}

func TestJWTManager_Verify_Invalid(t *testing.T) {
	userName, duration := random.NString(5), time.Minute
	p, err := token.NewPayload(userName, duration)
	require.NoError(t, err)
	jt := jwt.NewWithClaims(jwt.SigningMethodNone, p)
	token, err := jt.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)
	manager, err := New(random.NString(32))
	require.NoError(t, err)
	payload, err := manager.Verify(token)
	require.Error(t, err)
	require.Nil(t, payload)
}
