package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const secretkey string = "123456"

type CustomClaims struct {
	UserId   uint     `json:"userId"`
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
	jwt.StandardClaims
}

func GenerateToken(claims *CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretkey))
}

func JwtParse() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")

		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretkey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token invalid"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			c.Set("userId", claims.UserId)
			c.Set("roles", claims.Roles)
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}

	}
}
