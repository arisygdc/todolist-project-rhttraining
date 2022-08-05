package core

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/todolist-project-rhttraining/src/gateway/pkg/pb"
	"go-micro.dev/v4/client"
	"golang.org/x/net/context"
	"log"
	"strings"
)

func Trim(str string) string {
	return strings.Trim(str, " ")
}

func lengRule(min int, max int, str string) bool {
	return len(str) > max || len(str) < min
}

func IsAlphaNumeric(str string) bool {
	for _, v := range str {
		// Allowed lower alphabet and numeric
		if (v < 48 || (v > 57 && v < 97)) || v > 122 {
			return false
		}
	}
	return true
}

type AuthCheck interface {
	VerifyToken(ctx context.Context, in *pb.Session, opts ...client.CallOption) (*pb.IdResponse, error)
}

var ErrInvalidToken = fmt.Errorf("invalid token")

func validateToken(c echo.Context, checker AuthCheck) (string, error) {
	token, ok := c.Get(CTX_SESSION_TOKEN_KEY).(string)
	log.Printf("get token from context: %s\n", token)
	if !ok {
		return "", ErrInvalidToken
	}

	param := pb.Session{
		Token: token,
	}

	sesId, err := checker.VerifyToken(c.Request().Context(), &param)
	if err != nil {
		return "", err
	}

	emptyId := uuid.UUID{}
	if sesId.Id == emptyId.String() {
		return "", ErrInvalidToken
	}

	return sesId.Id, nil
}
