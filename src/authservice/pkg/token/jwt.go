package token

import (
	"fmt"
	jwtProvider "github.com/golang-jwt/jwt/v4"
)

var signingMethod = jwtProvider.SigningMethodHS256

type JWToken struct {
	secret string
}

func NewJWT(secret string) JWToken {
	return JWToken{
		secret: secret,
	}
}

func (jwt JWToken) Generate(payload Payload) (string, error) {
	claim := jwtProvider.NewWithClaims(signingMethod, payload)
	return claim.SignedString([]byte(jwt.secret))
}

// ParseWithClaimToken parse @signedToken to Payload
// @payloadClaim as pointer to claim payload struct
func (jwt JWToken) ParseWithClaimToken(signedToken string, payloadClaim Payload) error {
	keyFunc := func(t *jwtProvider.Token) (interface{}, error) {
		if t.Method.Alg() != signingMethod.Alg() {
			return nil, fmt.Errorf("signing method not provide")
		}
		return []byte(jwt.secret), nil
	}
	_, err := jwtProvider.ParseWithClaims(signedToken, jwtProvider.Claims(payloadClaim), keyFunc)
	return err
}
