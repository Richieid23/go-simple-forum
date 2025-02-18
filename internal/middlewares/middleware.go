package middlewares

import (
	"errors"
	"github.com/Richieid23/simple-forum/internal/configs"
	"github.com/Richieid23/simple-forum/pkg/jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey := configs.Get().Jwt.SecretKey

	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = strings.TrimSpace(header)
		if header == "" {
			c.AbortWithError(http.StatusUnauthorized, errors.New("Authorization header is empty"))
			return
		}

		userId, username, err := jwt.ValidateToken(header, secretKey)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		c.Set("userId", userId)
		c.Set("username", username)
		c.Next()
	}
}
