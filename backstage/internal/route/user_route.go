package route

import (
	"database/sql"
	"leaveManager/internal/identification/delivery/http"

	"github.com/gin-gonic/gin"
)

func NewLoginRoutern(db *sql.DB, group *gin.RouterGroup, groupjwt *gin.RouterGroup) {
	lc := http.NewLoginController(db)

	group.POST("/login", lc.Login)
	groupjwt.GET("/jwt", func(c *gin.Context) {
		c.String(200, "/jwt")
	})
	groupjwt.POST("/user", lc.CreateUser)

	groupjwt.GET("/user", lc.GetUserInfo)
}
