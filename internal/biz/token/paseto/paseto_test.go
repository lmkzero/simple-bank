// Package paseto 提供Paseto Token实现
package paseto

import (
	"testing"
	"time"

	"github.com/lmkzero/simple-bank/internal/random"
	"github.com/stretchr/testify/require"
)

func TestPasetoManager_Verify_No_Error(t *testing.T) {
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

func TestPasetoManager_Verify_Expired(t *testing.T) {
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
