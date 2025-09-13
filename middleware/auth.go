package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/haloapping/jejakmakan-api/api/user"
	"github.com/labstack/echo/v4"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// get Authorization header
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Missing Authorization header")
		}

		// expected: "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Authorization format")
		}
		tokenStr := parts[1]

		// parse and validate JWT
		jwtSecret := os.Getenv("JWT_SECRET_KEY")
		token, err := user.ParseToken(tokenStr, jwtSecret)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid or expired token")
		}

		// optional: set user claims to context
		c.Set("user", token)

		return next(c)
	}
}
