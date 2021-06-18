package middleware

import (
	"go-erp/application/handler"
	"go-erp/domain/throwable"
	"net/http"

	"github.com/gin-gonic/gin"
)

type securityMiddleware struct {
	handler.Handler
}

// Auth ...
func (s *securityMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokens := make(map[string]bool)

		// get authorization
		tokens["goDDD"] = true

		// token := c.Request.Header.Get("Authorization")

		// if _, exists := tokens[token]; exists {
		if _, exists := tokens["goDDD"]; exists {
			c.Next()
		} else {
			err := throwable.NewUnauthorized("Authorization required!")
			s.Error(c.Writer, err)
			c.AbortWithError(http.StatusInternalServerError, *err)

		}
	}
}

func NewSecurityMiddleware() *securityMiddleware {
	return &securityMiddleware{}
}
