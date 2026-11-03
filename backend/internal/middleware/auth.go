package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"

	"golearn/internal/response"
	"golearn/internal/service"
)

const UserIDKey = "user_id"

func Auth(tokens *service.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			response.Fail(c, 401, "unauthorized", "Token akses tidak ditemukan.")
			return
		}
		claims, err := tokens.Parse(strings.TrimPrefix(header, "Bearer "), "access")
		if err != nil {
			response.Fail(c, 401, "unauthorized", "Sesi tidak valid atau sudah berakhir.")
			return
		}
		c.Set(UserIDKey, claims.UserID)
		c.Next()
	}
}

// OptionalAuth attaches the user when a valid bearer token is present, and never blocks.
func OptionalAuth(tokens *service.TokenService) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if strings.HasPrefix(header, "Bearer ") {
			if claims, err := tokens.Parse(strings.TrimPrefix(header, "Bearer "), "access"); err == nil {
				c.Set(UserIDKey, claims.UserID)
			}
		}
		c.Next()
	}
}

func UserID(c *gin.Context) uint64 {
	v, ok := c.Get(UserIDKey)
	if !ok {
		return 0
	}
	id, _ := v.(uint64)
	return id
}
