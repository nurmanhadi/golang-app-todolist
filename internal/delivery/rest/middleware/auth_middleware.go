package middleware

import (
	"golang-app-todolist/internal/service"
	"golang-app-todolist/pkg/exception"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MiddlewareConfig struct {
	Viper *viper.Viper
	Log   *logrus.Logger
}

func (m *MiddlewareConfig) Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenString, err := getTokenFromHeader(c)
		if err != nil {
			m.Log.WithError(err).Warn("failed get token from header")
			return exception.NewError(401, "token is null")
		}
		jwt, err := service.JwtVerifyToken(tokenString, []byte(m.Viper.GetString("jwt.key")))
		if err != nil {
			m.Log.WithField("error", err).Warn("failed verify jwt token")
			return exception.NewError(401, "failed verify token")
		}
		c.Locals("username", jwt.Username)
		return c.Next()
	}
}

func getTokenFromHeader(c *fiber.Ctx) (string, error) {
	header := c.Get("Authorization")
	if header == "" {
		return "", exception.NewError(401, "null token authorization")
	}
	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", exception.NewError(401, "invalid token format")
	}
	return parts[1], nil
}
