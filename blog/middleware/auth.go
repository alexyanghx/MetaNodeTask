package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckPrevilege(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, ok := c.Get("roles")
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "没有权限"})
			return
		}

		for _, r := range roles.([]string) {
			if role == r {
				c.Next()
				return
			}
		}

		//没有权限
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "没有权限"})
	}
}
