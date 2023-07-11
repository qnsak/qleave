package middleware

import (
	"leaveManager/internal/jwtToken"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Header["Token"] != nil {
			authToken := c.Request.Header.Get("Token")

			authorized, err := jwtToken.IsAuthorized(authToken, secret)

			if err != nil {
				c.JSON(http.StatusUnauthorized, "no authHeader")
				c.Abort()
				return
			}

			if authorized {
				userID, err := jwtToken.ExtractIDFromToken(authToken, "golang")
				if err != nil {
					c.JSON(http.StatusUnauthorized, "authHeader")
					c.Abort()
					return
				}
				log.Println(userID)
				log.Println("----test---")
				c.Set("x-user-id", userID)
				c.Next()
				return
			}

			c.JSON(http.StatusUnauthorized, "no ")
			c.Abort()
			return
		} else {
			c.JSON(http.StatusUnauthorized, "error")
			c.Abort()
			return
		}
	}
}
