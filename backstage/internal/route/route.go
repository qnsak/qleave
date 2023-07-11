package route

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	middleware "leaveManager/internal/route/middleware"
)

func Setup(db *sql.DB, gin *gin.Engine) {
	publicRouter := gin.Group("")
	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	publicRouter.Use(CORSMiddleware())
	protectedRouter.Use(CORSMiddleware())
	protectedRouter.Use(middleware.JwtAuthMiddleware("golang"))

	NewLoginRoutern(db, publicRouter, protectedRouter)
	NewLeaveRoutern(db, publicRouter, protectedRouter)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
