package token

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Payload interface {
	Valid() error
}

type SessionPayload struct {
	Username string
	IssAt    int64
	ExpAt    int64
}

func NewSessionPayload(username string) SessionPayload {
	dur := 30 * time.Minute
	return SessionPayload{
		Username: username,
		IssAt:    time.Now().Unix(),
		ExpAt:    time.Now().Add(dur).Unix(),
	}
}
func (sp SessionPayload) Valid() error {
	if sp.ExpAt <= time.Now().Unix() {
		return jwt.ErrTokenExpired
	}
	return nil
}
