package auth

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestAuth(t *testing.T) {
	emailAuth := NewAuthenticator(EmailAuthMixin{})
	result, err := emailAuth.Authenticate("user@example.com", "password")
	assert.Equal(t, true, result)
	assert.Equal(t, nil, err)
}
