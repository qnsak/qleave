package route

import (
	"database/sql"
	"leaveManager/internal/leave/delivery/http"

	"github.com/gin-gonic/gin"
)

func NewLeaveRoutern(db *sql.DB, group *gin.RouterGroup, groupjwt *gin.RouterGroup) {
	lc, err := http.NewLeaveController(http.WithSqliteRepository(db))

	if err != nil {
		panic(err)
	}

	groupjwt.GET("/leave", lc.GetLeaveList)

	groupjwt.POST("/leave", lc.ApplyLeave)

	groupjwt.PUT("/leave", lc.ApproveLeave)

	groupjwt.DELETE("/leave/:id", lc.CancelLeave)

	groupjwt.GET("/leave/approve", lc.GetApproveList)

	groupjwt.GET("/attendance/:departmentId", lc.GetAttendanceRecord)

	groupjwt.GET("/leave/coming", lc.GetLeaveIsComing)

	groupjwt.GET("/leave/type", lc.GetLeaveType)

	groupjwt.GET("/leave/apply", lc.GetApplyLeaveInfo)
}
