package core

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strings"
)

const (
	AuthorizationProvideKey = "bearer"

	CTX_SESSION_TOKEN_KEY = "session_token"
)

func AuthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		log.Printf("middleware get authorization: %s\n", authHeader)

		httpErrCode := http.StatusUnauthorized
		if len(authHeader) < 110 {
			return echo.NewHTTPError(httpErrCode, http.StatusText(httpErrCode))
		}

		authorization := strings.Split(authHeader, " ")
		httpErrCode = http.StatusForbidden
		if len(authorization) != 2 {
			return echo.NewHTTPError(httpErrCode, http.StatusText(httpErrCode))
		}

		if strings.ToLower(authorization[0]) != AuthorizationProvideKey {
			return echo.NewHTTPError(httpErrCode, "authorization not provide")
		}

		token := authorization[1]
		log.Printf("middleware get Token: %s\n", token)
		splitToken := strings.Split(token, ".")
		if len(splitToken) != 3 {
			return echo.NewHTTPError(httpErrCode, http.StatusText(httpErrCode))
		}

		c.Set(CTX_SESSION_TOKEN_KEY, token)
		return next(c)
	}
}
