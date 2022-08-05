package token

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var secret = "8be78fda62b3842d8047e2396a0751e3aa7ef78bc23e1c694fe36dfca3080aa0"

func TestJWT(t *testing.T) {
	// Create auth payload
	username := "arisygdc"
	payload := NewSessionPayload(username)
	err := payload.Valid()
	assert.NoError(t, err)

	// Generate jwt signed string
	jwt := NewJWT(secret)
	signedString, err := jwt.Generate(&payload)
	assert.NoError(t, err)
	assert.NotEmpty(t, signedString)

	// Claim jwt
	var payloadClaim SessionPayload
	err = jwt.ParseWithClaimToken(signedString, &payloadClaim)
	assert.Equal(t, payload, payloadClaim)
	assert.NoError(t, err)
}
