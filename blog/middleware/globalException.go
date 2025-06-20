package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GlobalException() gin.HandlerFunc {

	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatusJSON(500, gin.H{"Msg": fmt.Sprintf("%v", err)})
			}

		}()
		c.Next()
	}

}
