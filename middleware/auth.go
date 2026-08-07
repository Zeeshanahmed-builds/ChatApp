package middleware

import (
	"github.com/Zeeshanahmed-builds/ChatApp/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(401, gin.H{"error": "token required"})
			c.Abort()
			return
		}
		tokenString := strings.Split(authHeader, "Bearer ")[1]
		token, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(400, gin.H{"error": "token is not valid "})

			c.Abort()
			return
		}

		// parse token claims

		// set the user id from claims into the context

		// parse token claims
		claims := token.Claims.(jwt.MapClaims)

		// set the user id from claims into the context
		userID := int(claims["user_id"].(float64))

		fmt.Println(userID)
		c.Set("userID", userID)

		c.Next()

	}

}
