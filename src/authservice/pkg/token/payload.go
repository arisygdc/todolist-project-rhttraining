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
	Exp      int64
}

func NewSessionPayload(username string) SessionPayload {
	dur := 30 * time.Minute
	return SessionPayload{
		Username: username,
		Exp:      time.Now().Add(dur).Unix(),
	}
}
func (sp SessionPayload) Valid() error {
	if sp.Exp <= time.Now().Unix() {
		return jwt.ErrTokenExpired
	}
	return nil
}
